package main

import (
	"github.com/ksubedi/go-web-seed/controllers"
	"github.com/ksubedi/go-web-seed/core"
	"github.com/ksubedi/go-web-seed/middlewares"
)

// getRouter returns the routers
func getRouter() (router *core.Router) {
	router = core.NewRouter()

	// All routes go here
	router.HandleFunc("/", controllers.HomePage)

	//Static Controller
	router.PathPrefix("/").Handler(&controllers.Static{"/static/public", router})

	// All middlewares go here
	router.AddMiddleware("/", &middlewares.HTTPLogger{})

	return
}
