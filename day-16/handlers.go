package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "home.page.tmpl".
	renderTemplates(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "about.page.html".
	renderTemplates(w, "about.page.tmpl")
}