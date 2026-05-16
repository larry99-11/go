package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// This var is a template cache to make accessing our templates quicker.
var templatesCache = make(map[string]*template.Template)

/*
function signature: renderTemplate(write http.ResponseWriter, tmpl string)

@Params: write -> takes a response writter

@Params: tmpl -> this takes a .tmpl file we want to parse and write to the browser window
*/
func RenderTemplate(write http.ResponseWriter, tmpl string) {
	// Parsing our template files
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//err := parsedTemplate.Execute(w, nil)

	// Writting our response writer to our template
	parsedTemplate.Execute(write, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
func createTempate(temp string) error {

	// This will store a slice of our templates location
	templates := []string{
		fmt.Sprintf("./templates/%s", temp),
		"./templates/base.layout.tmpl",
	}
	// Parse the templates from our slice
	//This will take each entry from our slice of strings and put them in as individual strings
	parse, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// Add our 'parse' template to our map
	templatesCache[temp] = parse
	return nil
}

func renderedTemplateCheck(write http.ResponseWriter, temp string) {
	// checking to see if we have any templates in our cache
	var tmpl *template.Template
	var err error

	// this would either set our inMap var to true or false, i.e true means the template is in the cache /false then template is not in the cache
	_, inMap := templatesCache[temp]

	// if its false create a template
	if !inMap {
		err = createTempate(temp)
		if err != nil {
			log.Println(err)
		}
	} else {
		// then we already have the cache in the template
		log.Println("Using cached template")
	}

	tmpl = templatesCache[temp]
	err = tmpl.Execute(write, nil)
	if err != nil {
		log.Println(err)
	}
}
