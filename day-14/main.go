package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
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

// Divide is the divide page handler 
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(12.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, "When 12 is divided by 6 the result is %f", f)
}

// divideValues divides two numbers and returns the result.
// If the divisor is zero, it returns an error.
func divideValues(x, y float32) (float32, error) {
	if y == 0.0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("starting the application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
