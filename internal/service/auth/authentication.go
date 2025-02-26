package auth

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"io"
	"net/http"
	"time"
)

const (
	ExpAT            = time.Minute * 10 //set based on Identity Provider expire Access Token
	ExpRT            = time.Hour * 24   // set based on Identity Provider expire Refresh Token
	ExpPk            = time.Minute * 10 //set based on Identity Provider expire Public Key
	KeycloakIssuer   = "http://localhost:8080/realms/qdrant-go-realm"
	KeycloakJWKSURL  = KeycloakIssuer + "/protocol/openid-connect/certs"
	KeycloakTokenURL = KeycloakIssuer + "/protocol/openid-connect/token"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OnlyAccessToken struct {
	AccessToken string
}

// token, Make HTTP request to get both tokens
func token(code string, conf *config.Config) (Token, error) {
	d := Director{builder: NewBuilder()}
	data := d.BuildAuthCodeData(conf, code, "authorization_code")
	resp, err := http.PostForm(KeycloakTokenURL, *data)
	if err != nil {
		return Token{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Token{}, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Token{}, err
	}

	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return Token{}, err
	}
	return t, nil

}

// NewAccessToken , used to get new access token with refresh token,
// if refresh expired return 401 to login again
func NewAccessToken(ctx context.Context, conf *config.Config, refreshToken string) (OnlyAccessToken, error) {
	d := Director{builder: NewBuilder()}
	data := d.BuildRefreshTokenData(conf, refreshToken, "refresh_token")
	resp, err := http.PostForm(KeycloakTokenURL, *data)
	if err != nil {
		return OnlyAccessToken{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return OnlyAccessToken{}, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OnlyAccessToken{}, err
	}

	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return OnlyAccessToken{}, err
	}
	return OnlyAccessToken{AccessToken: t.AccessToken}, nil
}

// TokenFromAuthCode , get the access and refresh token using authorization code.
// Return the access token and save the refresh toke to the cache
func TokenFromAuthCode(ctx context.Context, conf *config.Config, c repository.Cache, code string) (OnlyAccessToken, error) {
	t, err := token(code, conf)

	if err != nil {
		return OnlyAccessToken{}, err
	}

	claims, err := DecodeToken(ctx, conf, c, t.AccessToken)
	if err != nil {
		return OnlyAccessToken{}, err
	}
	// save refresh token to the cache
	err = c.Set(ctx, makeKey(claims["email"].(string), "_rk"), t.RefreshToken, ExpRT)
	if err != nil {
		return OnlyAccessToken{}, err
	}

	return OnlyAccessToken{
		AccessToken: t.AccessToken,
	}, nil

}

// Function to fetch JWKS from Keycloak
func fetchPK(token *jwt.Token) (interface{}, error) {

	// Fetch the JWKS
	resp, err := http.Get(KeycloakJWKSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKS response: %v", err)
	}

	// Parse JWKS
	keySet, err := jwk.ParseString(string(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWKS: %v", err)
	}

	// Find the correct key using the "kid" field
	if kid, ok := token.Header["kid"].(string); ok {
		key, found := keySet.LookupKeyID(kid)
		if !found {
			return nil, fmt.Errorf("public key with kid %s not found", kid)
		}

		var pubKey interface{}
		err = key.Raw(&pubKey)
		if err != nil {
			return nil, fmt.Errorf("failed to convert key: %v", err)
		}
		return pubKey, nil
	}
	return nil, fmt.Errorf("kid not found in token header")
}

// keyFunc, is higher order function cache public key. If not exist in cache it fetches from IP.
// Returns jwt.Keyfunc function type which is an argument type expected by the jwt.Parse function
func keyFunc(ctx context.Context, c repository.Cache) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {

		// Ensure the token uses RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		// Check PK exist in Cache if not  then fetch from keycloak
		pkByte, err := c.Get(ctx, makeKey(email, "_pk"))
		if err != nil {
			// fetch pk from the identity provider
			pk, err := fetchPK(token)
			if err != nil {
				return nil, err
			}
			// Convert rsa.Publickey to json byte
			jpk, err := rsaPublicKeyJSON(pk.(*rsa.PublicKey))
			if err != nil {
				return nil, err
			}

			//set into cache
			err = c.Set(ctx, makeKey(email, "_pk"), jpk, ExpPk)
			if err != nil {
				return nil, err
			}

			return pk, nil
		} else {
			pk, err := jsonRSAPublicKey(pkByte.(string))
			if err != nil {
				return nil, err
			}
			return pk, nil
		}

	}
}

// DecodeToken, decode and returns claims of token
func DecodeToken(ctx context.Context, conf *config.Config, c repository.Cache, tokenString string) (map[string]interface{}, error) {

	// Parse and validate the token
	t, err := jwt.Parse(tokenString, keyFunc(ctx, c))

	//we assume toke is expired, get new token with refresh token
	if err != nil {
		claims := t.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		rk, err := c.Get(ctx, makeKey(email, "_rk"))
		if err != nil {
			return nil, fmt.Errorf("Login again your session is ended: %v", err)
		}
		at, err := NewAccessToken(ctx, conf, rk.(string))
		if err != nil {
			return nil, err
		}
		return DecodeToken(ctx, conf, c, at.AccessToken)
	}

	// Check if the token is valid
	if !t.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Extract claims
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims")
	}

	// Validate `iss` (issuer)
	if claims["iss"] != KeycloakIssuer {
		return nil, fmt.Errorf("invalid issuer: %s", claims["iss"])
	}

	return claims, nil
}

// makeKey, build new key from given set of stings
func makeKey(k1 string, ks ...string) string {

	extra := ""
	for _, k := range ks {
		extra += k
	}
	return k1 + extra
}
