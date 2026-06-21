package main

/*
TODO:
- Have a route for /
-  Have a custom route for i.e /main or /login etc...
*/
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/larry-99/webGo/pkg/config"
	"github.com/larry-99/webGo/pkg/handlers"
	"github.com/larry-99/webGo/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

/*
http.HandlerFunc to convert simple functions into handlers.
The function http.HandleFunc to tell the server which function to call to handle a request to the server.
*/
func main() {
	// change this to true when in production
	app.InProd = false

	// Initialize a new session manager and configure the session lifetime.
	sessionCookie := scs.New()
	/*
		Lifetime controls the maximum length of time that a session is valid for before it expires.
		 The default value is 24 hours.
	*/
	sessionCookie.Lifetime = 24 * time.Hour

	// setting params for our cookie
	sessionCookie.Cookie.Persist = true
	sessionCookie.Cookie.SameSite = http.SameSiteLaxMode
	sessionCookie.Cookie.Secure = app.InProd // only for dev not prod

	app.Session = sessionCookie
	tc, err := render.CreateTempateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting server! Listening on port %s", portNumber)

	serve := &http.Server{
		// Addr optionally specifies the TCP address for the server to listen on,
		// in the form "host:port". If empty, ":http" (port 80) is used.
		Addr: portNumber,

		// This will handle our routes deinfed in routes.go
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
