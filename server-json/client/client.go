package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	getUsers()
	postUser()
}

func getUsers() {
	url := "http://localhost:7575/user"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received response: %s\n", string(body))

}

func postUser() {
	url := "http://localhost:7575/user"
	data := []byte(`
	{
		"name": "bar",
		"age": 4
	}
	`)

	r := bytes.NewReader(data)

	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received response: %s\n", string(body))
}
