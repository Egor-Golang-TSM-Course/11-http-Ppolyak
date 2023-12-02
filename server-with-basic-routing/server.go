package main

import (
	"fmt"
	"net/http"
)

import (
	time2 "time"
)

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
git