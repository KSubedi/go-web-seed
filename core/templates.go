package core

import (
	"fmt"
	"html/template"
	"net/http"
)

// Directory that holds all the template files
const tmplDirectory string = "templates/"

// TemplateData is a struct that holds a map that holds all the data that is going to be passed to templates
type TemplateData struct {
	Data map[string]string
}

// DoesExist is a function available to the templates that allows you to
func (t *TemplateData) DoesExist(field string) bool {
	_, ok := t.Data[field]
	if ok {
		return true
	}
	return false
}

// RunTemplate will run a template and output to the io writer passed as third parameter
func RunTemplate(name string, data interface{}, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(tmplDirectory + name + ".html")
	if err != nil {
		fmt.Println("Error Occured!")
		fmt.Println(err)
	}
	tmpl.Execute(w, data)
}
