package main

import (
	"fmt"
	"github.com/johnarumemi/go-course/pkg/config"
	"github.com/johnarumemi/go-course/pkg/handlers"
	"github.com/johnarumemi/go-course/pkg/render"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	var app config.AppConfig

	// Set Template Cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// Create and Set Repo
	repo := handlers.CreateRepo(&app)

	handlers.SetRepo(repo)

	render.NewTemplates(&app)

	// Request must be a pointer, since the Request is passed to us
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port:%d\n", PORT)

	// start web server and listen on port; ignore the error from it.
	_ = http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	// ports are just things you can listen to from applications
}
