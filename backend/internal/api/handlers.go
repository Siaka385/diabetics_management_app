package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	// auth "diawise/internal/auth"

	"gorm.io/gorm"
)

/*
*	Frontend server
*
 */
func Index(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "index.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GlucoseTrackerEndPointHandler(w http.ResponseWriter, r *http.Request) {
	// Capture glucose level and date from the request query parameters
	glucoseLevel := r.URL.Query().Get("glucose")
	glucoseDate := r.URL.Query().Get("date")

	glucoseParam := map[string]string{glucoseLevel: glucoseDate}

	// Set response header and JSON encode the glucose level and date
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(glucoseParam)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, ok := Posts[postID]
	if !ok {
		NotFoundHandler(w)
		return
	}

	tmpl, err := template.ParseFiles(
		"../frontend/public/base.html",
		"../frontend/public/blog_display.html",
	)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", post)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}
}

func BlogHomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"../frontend/public/base.html",
		"../frontend/public/blog_home.html",
	)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", Data); err != nil {
		InternalServerErrorHandler(w)
		return
	}
}

func BadRequestHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusBadRequest
	hitch.Problem = "Bad Request!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func InternalServerErrorHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusInternalServerError
	hitch.Problem = "Internal Server Error!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func NotFoundHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusNotFound
	hitch.Problem = "Not Found!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}
