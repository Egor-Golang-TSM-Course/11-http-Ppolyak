package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getHomePage() {

	homeUrl := "http://localhost:7575/"
	//timeUrl := "http://localhost:7575/time"
	response, err := http.Get(homeUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//response.Body.Close()
	fmt.Printf("Received response: %s\n", string(body))

}
