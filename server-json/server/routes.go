package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.MethodFunc("Post", "/user", postHandler)

	return r
}
