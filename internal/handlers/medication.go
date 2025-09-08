package handlers

import (
	"html/template"
	"net/http"
)

func MedicationHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "medication.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}
