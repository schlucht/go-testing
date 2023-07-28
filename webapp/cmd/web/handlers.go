package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	ip := r.URL.Path

	err := app.render(w, r, "home.page", &TemplateData{
		IP: ip,
	})

	if err != nil {
		log.Fatal(err)
	}

}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk
	parsedTemplate, err := template.ParseFiles(("./templates/" + t + ".gohtml"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	// executed tamplate
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
