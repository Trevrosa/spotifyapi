package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	token, err := getAccessToken(client)
	if err != nil {
		panic("we couldnt get the access token: " + err.Error())
	}

	fmt.Println(token.AccessToken)
}
