package handlers

import (
	"net/http"

	"github.com/adiputra22/learn-golang-web/pkg/config"
	"github.com/adiputra22/learn-golang-web/pkg/models"
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
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["person"] = "Adiputra"

	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
