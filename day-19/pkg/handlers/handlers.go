package handlers

import (
	"net/http"
	"github.com/Learn-Golang/modules/day-19/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "home.page.tmpl".
	render.RenderTemplates(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "about.page.html".
	render.RenderTemplates(w, "about.page.tmpl")
}