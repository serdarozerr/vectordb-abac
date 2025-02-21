package model

import "context"

type AuthCode struct {
	Code string `json:"authorization_code"`
}

func (auth AuthCode) Valid(ctx context.Context) (problems map[string]string) {

	problems = map[string]string{}

	if auth.Code == "" {
		problems["Code Error"] = "Authorization Code is missing"
	}
	return problems

}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

func (auth AccessToken) Valid(ctx context.Context) (problems map[string]string) {

	problems = map[string]string{}

	if auth.AccessToken == "" {
		problems["Token Error"] = "Access Token is missing"
	}
	return problems

}
