package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int `json:"expires_in"`
}

// valid for 1 hour
func getAccessToken(client *http.Client) (tokenResponse *TokenResponse, err error) {
	const API_URL = "https://accounts.spotify.com/api/token"

	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	clientId, ok := env["CLIENT_ID"]
	if !ok {
		return nil, errors.New("no CLIENT_ID environment variable")
	}
	clientSecret, ok := env["CLIENT_SECRET"]
	if !ok {
		return nil, errors.New("no CLIENT_SECRET environment variable")
	}

	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", clientId)
	params.Set("client_secret", clientSecret)

	resp, err := client.PostForm(API_URL, params)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("request error " + resp.Status + ": " + string(body))
	}


	token := TokenResponse{}
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
