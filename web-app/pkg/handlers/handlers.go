package handlers

import (
	"net/http"

	"github.com/larry-99/webGo/pkg/config"
	"github.com/larry-99/webGo/pkg/models"
	"github.com/larry-99/webGo/pkg/render"
)

// Repo: The repository used by the handlers
var Repo *Repository

// Repository: Is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo: This function creates a new repository
func NewRepo(conf *config.AppConfig) *Repository {
	return &Repository{
		App: conf,
	}
}

// NewHandlers: Sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (r *Repository) Home(write http.ResponseWriter, req *http.Request) {
	// Grab the remote IP addr of user visiting the site and store it as a K/V pair
	remoteIp := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIp)

	stringMap := make(map[string]string)
	stringMap["aTest"] = "test"
	tmplName := "home.page.tmpl"

	render.RenderTemplate(write, tmplName, &models.TemplateData{
		StringMap: stringMap,
	})

}

//to pull the IP addr
// remoteIP := r.App.Session.GetString(r.context(), "remote_ip")
