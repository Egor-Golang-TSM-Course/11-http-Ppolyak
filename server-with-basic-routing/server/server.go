package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	startServer()
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/time", showTime)
	log.Println("Запуск веб-сервера на http://127.0.0.1:7575")
	err := http.ListenAndServe(":7575", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Println("Home page opened")
	w.Write([]byte("Home page"))
}

func showTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.NotFound(w, r)
		return
	}
	log.Println("Time page opened")
	ss := fmt.Sprintf("Time: %s\n", time.Now().String())
	w.Write([]byte(ss))
}
