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
	"net/url"
	"time"
)

const (
	ExpPk           = time.Minute * 10 //expire in 10 min.
	KeycloakIssuer  = "http://localhost:8080/realms/qdrant-go-realm"
	KeycloakJWKSURL = KeycloakIssuer + "/protocol/openid-connect/certs"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OnlyAccessToken struct {
	AccessToken string
}

// getToken, Make HTTP request to provide to get tokens
func getToken(code string, conf *config.Config) (Token, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", conf.Auth.GrantType)
	data.Set("redirect_uri", conf.Auth.RedirectURI)
	data.Set("client_id", conf.Auth.ClientID)
	data.Set("client_secret", conf.Auth.ClientSecret)
	resp, err := http.PostForm(conf.Auth.TokenURL, data)
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
	var token Token

	err = json.Unmarshal(body, &token)
	if err != nil {
		return Token{}, err
	}
	return token, nil

}

// RefreshTheToken , used to get new access token with refresh token,
// if refresh expired return 401 to login again
func RefreshTheToken(token Token, conf *config.Config) error {
	return nil
}

// TokenFromCode , get the access and refresh token using authorization code.
// Return the access token and save the refresh toke to the redis db
func TokenFromCode(conf *config.Config, code string) (OnlyAccessToken, error) {
	token, err := getToken(code, conf)
	if err != nil {
		return OnlyAccessToken{}, err
	}
	// save refresh token to the Redis
	return OnlyAccessToken{
		AccessToken: token.AccessToken,
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

// parseToken, is higher order function cache public key. If not exist in cache it fetches from IP.
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
		pkByte, err := c.Get(ctx, email)
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
			err = c.Set(ctx, email, jpk, ExpPk)
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

// Function to decode and verify a Keycloak token
func DecodeToken(ctx context.Context, c repository.Cache, tokenString string) (bool, error) {

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, keyFunc(ctx, c))
	if err != nil {
		return false, fmt.Errorf("invalid token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return false, fmt.Errorf("token is not valid")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("failed to extract claims")
	}

	// Validate `iss` (issuer)
	if claims["iss"] != KeycloakIssuer {
		return false, fmt.Errorf("invalid issuer: %s", claims["iss"])
	}

	// Token is valid
	return token.Valid, nil
}
