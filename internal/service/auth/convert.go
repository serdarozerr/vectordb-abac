package auth

import (
	"crypto/rsa"
	"encoding/json"
	"math/big"
)

type rsaPublicKeyJson struct {
	N *big.Int
	E int
}

func rsaPublicKeyJSON(pk *rsa.PublicKey) ([]byte, error) {

	r := rsaPublicKeyJson{E: pk.E, N: pk.N}
	v, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return v, nil

}

func jsonRSAPublicKey(value string) (*rsa.PublicKey, error) {
	r := &rsaPublicKeyJson{}
	err := json.Unmarshal([]byte(value), r)
	if err != nil {
		return nil, err
	}
	return &rsa.PublicKey{E: r.E, N: r.N}, nil

}
