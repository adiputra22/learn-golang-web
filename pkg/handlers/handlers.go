package handlers

import (
	"net/http"

	"github.com/adiputra22/learn-golang-web/pkg/config"
	"github.com/adiputra22/learn-golang-web/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
