package handlers

import (
	"net/http"

	"github.com/Learn-Golang/modules/day-21/pkg/config"
	"github.com/Learn-Golang/modules/day-21/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repositotry is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "home.page.tmpl".
	render.RenderTemplates(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// The renderTemplates function is called with the value of tmpl as "about.page.html".
	render.RenderTemplates(w, "about.page.tmpl")
}