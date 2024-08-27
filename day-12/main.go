package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	Register a new handler function for the root URL path "/"
	This function will be executed whenever a request is made to "/"
	'w' is the http.ResponseWriter interface, used to send a response back to the client
	'r' is a pointer to an http.Request, containing all the information about the client's request
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "Hello World !!")

		if err != nil {
			fmt.Println(err)
		}
		/*
		The information below is printed on the terminal every time a
		request is made. Specifically, it prints the number of bytes
		written to the response.
		*/
		fmt.Printf("the no. of bytes are:%d\n", n)
	})
	/*
	Start the HTTP server and listen on port 8080 for incoming requests.
	The first argument ":8080" specifies that the server should listen 
	on port 8080 on all available network interfaces (IP addresses).
	The second argument 'nil' means we're using the default multiplexer 
	(http.DefaultServeMux), which routes incoming requests to the 
	appropriate handler functions.
	*/
	http.ListenAndServe(":8080", nil)
	// go to "localhost:8080" on your web browser
}
