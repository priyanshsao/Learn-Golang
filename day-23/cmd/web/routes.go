package main

import (
	"net/http"

	"github.com/Learn-Golang/modules/day-23/pkg/config"
	"github.com/Learn-Golang/modules/day-23/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {

	_ = app // This will eliminate the "app is unused" warning
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
