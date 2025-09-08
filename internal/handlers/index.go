package handlers

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

func Index(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templateName := "index.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}
}
