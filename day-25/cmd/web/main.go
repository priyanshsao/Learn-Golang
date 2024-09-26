package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Learn-Golang/modules/day-25/pkg/config"
	"github.com/Learn-Golang/modules/day-25/pkg/handlers"
	"github.com/Learn-Golang/modules/day-25/pkg/render"
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

	fmt.Printf("starting the application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
