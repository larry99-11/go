package render

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
function signature: renderTemplate(write http.ResponseWriter, tmpl string)

@Params: write -> takes a response writter

@Params: tmpl -> this takes a .tmpl file we want to parse and write to the browser window
*/
func RenderTemplate(write http.ResponseWriter, tmpl string) {
	// Parsing our template files
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	//err := parsedTemplate.Execute(w, nil)

	// Writting our response writer to our template
	parsedTemplate.Execute(write, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
