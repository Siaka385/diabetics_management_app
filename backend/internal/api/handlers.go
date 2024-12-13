package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type BlogPost struct {
	Title   string
	Author  string
	Date    string
	Content string
}

type Post struct {
	ID      string
	Title   string
	Excerpt string
}

type Issue struct {
	StatusCode int
	Problem    string
}

// Initialize variable to hold error message and status codes

var hitch Issue

var LoadTemplate = func() (*template.Template, error) {
	return template.ParseFiles("frontend/public/error.html")
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
	tmpl, err := template.ParseFiles(
		"frontend/public/blog_base.html",
		"frontend/public/blog_display.html",
	)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	data := BlogPost{
		Title:   "First Blog Post",
		Author:  "John Doe",
		Date:    "December 13, 2024",
		Content: "<p>This is the full content of the post.</p>",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}
}

func BlogHomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"frontend/public/blog_base.html",
		"frontend/public/blog_home.html",
	)

	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	data := struct {
		Title string
		Posts []Post
	}{
		Title: "Home",
		Posts: []Post{
			{ID: "1", Title: "First Blog Post", Excerpt: "This is the first post."},
			{ID: "2", Title: "Another Post", Excerpt: "Learn more about this topic."},
		},
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
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
