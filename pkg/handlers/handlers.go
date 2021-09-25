package handlers

import (
	"fmt"
	"github.com/johnarumemi/go-course/pkg/config"
	"github.com/johnarumemi/go-course/pkg/render"
	"log"
	"net/http"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap	map[string]string
	IntMap		map[string]int
	FloatMap	map[string]float32
	Data        map[string]interface{}
	CSRFToken	string
	Flash		string
	Warning		string
	Error 		string
}

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// CreateRepo creates a new repository
func CreateRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetRepo sets the repository for the handlers package
func SetRepo(r *Repository) {
	Repo = r
}

func RawHome(w http.ResponseWriter, r *http.Request) {
	// standard practice is to use 'w' and 'r' respectively
	_, err := fmt.Fprintf(w, "Hello world")

	if err != nil {
		log.Println("Error has occurred", err)
	}

}

// Home - Handler for homepage
// function that handlers a request
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some business logic

	// send data to the template
	render.RenderTemplate(w, "home.page.gohtml", &TemplateData{})
}

// About is the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &TemplateData{})
}
