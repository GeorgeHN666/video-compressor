package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/insert", StorePost)
	mux.Get("/getAll", Feed)

	return mux
}
