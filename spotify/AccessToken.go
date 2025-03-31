package spotify

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

// the response given by the api containing the access token
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// return the access token as a [http.Header]. 
func (resp *AccessTokenResponse) Header() http.Header {
	value := resp.TokenType + " " + resp.AccessToken
	return http.Header{"Authorization": []string{value}}
}

// valid for 1 hour
func GetAccessToken(client *http.Client) (*AccessTokenResponse, error) {
	const API_URL = "https://accounts.spotify.com/api/token"

	clientId, ok := os.LookupEnv("CLIENT_ID")
	if !ok {
		return nil, errors.New("no CLIENT_ID environment variable")
	}
	clientSecret, ok := os.LookupEnv("CLIENT_SECRET")
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

	token := AccessTokenResponse{}
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
