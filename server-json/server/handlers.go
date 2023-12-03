package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	var resp Person
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		err = json.Unmarshal(body, &resp)
		sl = append(sl, string(body))
	}
	if err != nil {
		fmt.Println(w, "can't unmarshal: ", err.Error())
	}

	fmt.Fprintf(w, "Name: %s, Age: %v", resp.Name, resp.Age)
	log.Println("Post request received")
	fmt.Println(sl)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}
	log.Println("Get request received")
	fmt.Fprintf(w, "Users list %s, size: %v", sl, len(sl))
	fmt.Println(sl)
}
