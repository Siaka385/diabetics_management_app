package api

import (
	"net/http"
	"html/template"

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
