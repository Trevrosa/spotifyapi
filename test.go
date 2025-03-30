package main

import (
	"fmt"
	"io"
	"net/http"
)

func car() string {
	return "I am a car"
}

// try pointers

func fakeAddOne(v int) {
	v += 1
}

func addOne(v *int) {
	*v += 1
}

// try http & try errors
func getIp(client *http.Client) (ip string, err error) {
	fmt.Println("getting api")
	resp, err := client.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("reading body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// try generics and arrays
func addOneEle[T any](arr []T, ele T) []T {
	return append(arr, ele)
}

func run() {
	fmt.Println(car())

	var v = 1
	fakeAddOne(v)
	addOne(&v)
	fmt.Println("1 + 1 is", v)

	var client = &http.Client{}

	ip, err := getIp(client)
	if err != nil {
		panic("failed to get ip: " + err.Error())
	}

	fmt.Printf("my ip is %s\n", ip)

	var list []string
	list = addOneEle(list, "one")
	list = addOneEle(list, "two")

	var list2 []string
	list2 = addOneEle(list2, "three")
	list2 = addOneEle(list2, "four")

	var lists = [][]string{list, list2}
	for i, list := range lists {
		fmt.Printf("list%d: %v\n", i+1, list)
	}
}
