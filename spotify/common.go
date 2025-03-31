package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SpotifyType interface {
	json.Marshaler
	SpotifyType()
}

type SpotifyError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err *SpotifyError) Error() string {
	return fmt.Sprintf("error %d: %s", err.Status, err.Message)
}

func ToSpotifyError(resp []byte) error {
	type errorResponse struct {
		Error SpotifyError `json:"error"`
	}

	err := errorResponse{}
	if err := json.Unmarshal(resp, &err); err != nil {
		return errors.New("unmarshal error: " + err.Error() + "\n" + string(resp))
	}
	return &err.Error
}

func GetAuthed(client *http.Client, url string, auth http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header = auth

	return client.Do(req)
}

func GetFormAuthed(client *http.Client, url string, params map[string]string, auth http.Header) (*http.Response, error) {
	parsedParams := ""
	for k, v := range params {
		parsedParams += fmt.Sprintf("%s=%s,", k, v)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", url, parsedParams), http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header = auth
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return client.Do(req)
}
