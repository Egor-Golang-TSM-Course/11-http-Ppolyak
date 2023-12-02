package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

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
