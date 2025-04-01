package main

import (
	"fmt"
	"net/http"
	"trev/spot/v2/spotify"
	"trev/spot/v2/spotify/artist"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	client := &http.Client{}

	if err := godotenv.Load(); err != nil {
		panic("couldnt load .env file: " + err.Error())
	}

	fmt.Println("getting access token from app credentials..")
	token, err := spotify.GetAccessToken(client)
	if err != nil {
		panic("we couldnt get the access token: " + err.Error())
	}

	auth := token.Header()

	fmt.Printf("\ngot token: ")
	spew.Dump(auth)
	fmt.Println()

	var artistId string
	fmt.Print("artist id: ")
	fmt.Scanln(&artistId)
	fmt.Println()

	artist, err := artists.GetArtist(client, artistId, auth)
	if err != nil {
		panic("artist fail: " + err.Error())
	}
	spew.Dump(artist)

	albums, err := artists.GetArtistAlbums(client, artistId, auth)
	if err != nil {
		panic("albums fail: " + err.Error())
	}

	fmt.Println("artists:")
	for _, album := range *albums {
		fmt.Println(album.Name)
	}
}
