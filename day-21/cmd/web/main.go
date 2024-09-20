package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Learn-Golang/modules/day-21/pkg/config"
	"github.com/Learn-Golang/modules/day-21/pkg/handlers"
	"github.com/Learn-Golang/modules/day-21/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting the application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
