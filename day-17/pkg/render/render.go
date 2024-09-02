package render

import (
	"html/template"
	"net/http"
	"fmt"
)


func RenderTemplates(w http.ResponseWriter, tmpl string) {
	/*
	   ParseFiles loads and parses the specified template file from the
	   given path and returns a parsed template object and an error.
	*/
	parsedTemplate, err := template.ParseFiles("D:/golang series/Learn-Golang/day-17/templates/" + tmpl)
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