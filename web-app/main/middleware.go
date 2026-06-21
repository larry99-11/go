package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func writeToConsole(next http.Handler) http.Handler {
	// adding a anonymous func in here to cast to the http.HandlerFunc()
	return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		fmt.Println("Hi there! You hit the page")
		// Move on to the next part
		next.ServeHTTP(write, req)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// the csfrs handler needs a base cookie to make sure that the token it generates is a available on a per page basis
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProd,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Automatically loads and saves session data for the current request, and communincates the session token
// to and from the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
