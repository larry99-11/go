package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/larry-99/webGo/pkg/config"
	"github.com/larry-99/webGo/pkg/handlers"
)

// creates a pattern mux (multiplexer)
func routes(app *config.AppConfig) http.Handler {

	mux := pat.New()
	// Where to find the static images
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	//Get will register a pattern with a handler for GET requests. This is routing to our home page.
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//chain the middleware together so every request passes through both — SessionLoad first, then NoSurf
	return SessionLoad(NoSurf(mux))
}

