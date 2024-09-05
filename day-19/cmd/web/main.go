package main

import (
	"fmt"
	"net/http"
	"github.com/Learn-Golang/modules/day-19/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting the application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
