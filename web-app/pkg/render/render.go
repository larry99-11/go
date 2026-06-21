package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/larry-99/webGo/pkg/config"
	"github.com/larry-99/webGo/pkg/models"
)

// This var is a template cache to make accessing our templates quicker.
var templatesCache = make(map[string]*template.Template)
var app *config.AppConfig


func AddDefaultData(tempData *models.TemplateData) *models.TemplateData {
	return tempData
}

// NewTemplates() sets the config for the template package
func NewTemplates(conf *config.AppConfig) {
	app = conf
}

/*
function signature: renderTemplate(write http.ResponseWriter, tmpl string)

@Params: write -> takes a response writter

@Params: tmpl -> this takes a .tmpl file we want to parse and write to the browser window
*/
func RenderTemplate(write http.ResponseWriter, tmpl string, tempData *models.TemplateData) {
	// get the template cache from the app config struct
	tmplCache := app.TemplateCache

	// get requested template from cache
	temp, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("Could not get template ")
	}

	// Return a default page
	tempData = AddDefaultData(tempData)

	// render the template
	temp.Execute(write, tempData)
}

func CreateTempateCache() (map[string]*template.Template, error) {

	// no need to use make(), just use {} when creating maps
	myCache := map[string]*template.Template{}

	// get all the files names *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range theough the pages since is returns a slice of strings
	for _, page := range pages {
		//page just holds the full path
		name := filepath.Base(page)

		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// looking for layouts.tmpl
		layoutMatches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// deal with the templates if there is anything inside them
		if len(layoutMatches) > 0 {
			// parseGlob() will parse the template definition
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet

	}

	return myCache, nil

}
