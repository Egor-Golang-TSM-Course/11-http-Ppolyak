package main

import (
	"log"
	"net/http"
)

func main() {
	// Обработчик HTTP-запросов
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/time", showTime)
	log.Println("Запуск веб-сервера на http://127.0.0.1:7575")
	err := http.ListenAndServe(":7575", mux)
	log.Fatal(err)

}
