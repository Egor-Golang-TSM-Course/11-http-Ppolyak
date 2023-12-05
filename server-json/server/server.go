package main

import (
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	log.Println("Запуск веб-сервера на http://127.0.0.1:7575")
	err := http.ListenAndServe(":7575", routes())
	log.Fatal(err)
}
