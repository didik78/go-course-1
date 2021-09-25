package render

import (
	"bytes"
	"errors"
	"github.com/johnarumemi/go-course/pkg/config"
	"github.com/johnarumemi/go-course/pkg/handlers"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// functions: A map of functions that can be used in a template, i.e. create our own template functions
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate writes template to response writer
// w: Response writer
// tmpl: Filepath of template to be parsed
func RenderTemplate(w http.ResponseWriter, tmpl string, td *handlers.TemplateData) {

	var tc map[string]*template.Template

	// get template cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	// get template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	// write template to buffer
	_ = t.Execute(buf, nil)

	// write to response writter
	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Create map of template names: *Template
	myCache := map[string]*template.Template{}
	/*
		about.page.gohtml: *Template
	*/

	// get all pages in template folder that have string "page" in them
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	// Iterate through pages
	for _, page := range pages {
		// page is full path to the page
		// get name of page only
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		// does this template match any layouts?

		// first find all layouts
		matches, _ := filepath.Glob("./templates/*.layout.gohtml")

		if len(matches) > 0 {
			/*
				parses the template definitions in the files identified by the pattern
				and associates the resulting templates with t.
			*/
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}
		}

		// [name of template/page] = *Template
		myCache[name] = ts
	}

	return myCache, nil
}

func AddValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return sum, nil
}
