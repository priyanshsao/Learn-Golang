package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
// ignore this function
func RenderTemplatesTest(w http.ResponseWriter, tmpl string) {
	/*
	   ParseFiles loads and parses the specified template file from the
	   given path and returns a parsed template object and an error.
	*/
	parsedTemplate, err := template.ParseFiles("D:/golang series/Learn-Golang/day-19/templates/"+tmpl, "D:/golang series/Learn-Golang/day-19/templates/base.layout.tmpl")
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

var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("D:/golang series/Learn-Golang/day-19/templates/%s", t),
		"D:/golang series/Learn-Golang/day-19/templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
