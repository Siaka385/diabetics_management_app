package api

import (
	"html/template"
	"net/http"

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

func BloodSugar(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "bloodsugar.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func DietAndNutrient(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "NutritionAndDiet.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Medicine(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "Medication.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Community(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "CommunityAndSupport.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Education(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		// auth.LoginUser(db, "toni", "antony102")

		var templateName string
		templateName = "Education.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
