package controllers

import (
	"net/http"

	"github.com/ksubedi/go-web-seed/core"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	//Make a new map to store key values passed to the template
	data := make(map[string]string)

	//Add data to the map
	data["Title"] = "Hello World"

	//Run a template from the templates directory and pass data to it
	core.RunTemplate("pages/index.html", core.TemplateData{data}, w)
}
