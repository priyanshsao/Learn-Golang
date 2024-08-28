package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(4, 2)
	fmt.Fprintf(w, "The value of x + y is %d", sum)
}

// addValue adds two integers and returns the sum
func addValue(x, y int) int {
	return x + y
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("starting the application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
