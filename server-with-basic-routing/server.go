package main

import (
	"fmt"
	"log"
	"net/http"
)

import (
	time2 "time"
)

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/time", showTime)
	log.Println("Запуск веб-сервера на http://127.0.0.1:7575")
	err := http.ListenAndServe(":7575", mux)
	getHomePage()
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Home page"))
}

func showTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.NotFound(w, r)
		return
	}
	time := time2.Now()
	ss := fmt.Sprintf("Time: %s\n", time)
	w.Write([]byte(ss))
}
