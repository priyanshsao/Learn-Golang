package handlers

import (
	"net/http"
	"github.com/Learn-Golang/modules/day-23/pkg/config"
	"github.com/Learn-Golang/modules/day-23/pkg/models"
	"github.com/Learn-Golang/modules/day-23/pkg/render"
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
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "(TEXT which is PASSED) "

	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
