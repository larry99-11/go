package handlers

import (
	"net/http"

	"github.com/larry-99/webGo/pkg/render"
)

func Home(write http.ResponseWriter, req *http.Request) {
	tmplName := "home.page.tmpl"
	render.RenderTemplate(write, tmplName)
}
