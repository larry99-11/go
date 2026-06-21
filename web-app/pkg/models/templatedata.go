package models

// TemplateData: holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}

	/* 
	cross sight request forgery token: When you build a web page you have a hidden field in that form, 
	which is a long string of random numbers and changes everytime someone goes to a page.
	*/
	CSRFToken string
	FlashMsg  string
	Warning   string
	Error     string
}