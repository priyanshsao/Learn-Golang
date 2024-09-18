package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	// creating a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	/* 
	   Create a buffer to temporarily hold the template output before
	   sending to the browser 
	*/
	buf := new(bytes.Buffer)
	
	// Execute the template and write its output to the buffer
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	// Write the buffer's content (HTML) to the HTTP response writer (the browser)
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./day-20/templates
	pages, err := filepath.Glob("D:/golang series/Learn-Golang/day-20/templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matching, err := filepath.Glob("D:/golang series/Learn-Golang/day-20/templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		/*
		   If there are any matching layout templates, parse them
		   and associate them with the page template
		*/
		if len(matching) > 0 {
			ts, err = ts.ParseGlob("D:/golang series/Learn-Golang/day-20/templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		/*
		  Store the fully parsed template (with both page and layout)
		  in the cache, using the template name as the key
		*/
		myCache[name] = ts
	}
	return myCache, nil
}
