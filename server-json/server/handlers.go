package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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
		if err != nil {
			fmt.Println(w, "can't unmarshal: ", err.Error())
		}
	}

	fmt.Fprintf(w, "Name: %s, Age: %v", resp.Name, resp.Age)

	log.Println("Post request received")
}
