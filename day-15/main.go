package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "home.page.tmpl".
	renderTemplates(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "about.page.html".
	renderTemplates(w, "about.page.tmpl")
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	/*
	   ParseFiles loads and parses the specified template file from the
	   given path and returns a parsed template object and an error.
	*/
	parsedTemplate, err := template.ParseFiles("D:/golang series/Learn-Golang/day-15/templates/" + tmpl)
	if err != nil {
		// http.Error writes an error message to the web page for users.
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		// fmt.Println prints the error message to the terminal for developers.
		fmt.Println("Error parsing template:", err)
		return
	}
	// Execute applies the parsed template to the response writer and sends it to the client.
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("starting the application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
