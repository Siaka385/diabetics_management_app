package handlers

import (
	"html/template"
	"net/http"
)

// TemplateData holds common template data
type TemplateData struct {
	Title string
	CSS   string
	JS    string
	User  interface{}
	Data  interface{}
}

// RenderProtectedTemplate renders a template with auth protection
func RenderProtectedTemplate(w http.ResponseWriter, tmpl *template.Template, templateName string, data TemplateData) {
	// Add auth scripts to protected pages
	if data.JS == "" {
		data.JS = "/static/js/protected-page.js"
	}

	err := tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		InternalServerErrorHandler(w)
	}
}

// RenderPublicTemplate renders a template for public pages
func RenderPublicTemplate(w http.ResponseWriter, tmpl *template.Template, templateName string, data TemplateData) {
	err := tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		InternalServerErrorHandler(w)
	}
}

// InternalServerErrorHandler handles internal server errors
func InternalServerErrorHandler(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
