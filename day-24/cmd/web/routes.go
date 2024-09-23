package main

import (
	"net/http"

	"github.com/Learn-Golang/modules/day-24/pkg/config"
	"github.com/Learn-Golang/modules/day-24/pkg/handlers"
	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	_ = app 

	mux := chi.NewRouter()
	
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
