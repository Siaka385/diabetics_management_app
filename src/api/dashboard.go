package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"diawise/src/auth"

	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		/* ------------------ HERE -----------------------------*/
		// The browser will now send cookies automatically once a user is logged in
		// you can use the datails from the user that is parsed from the cookie e.g user.ID (see line 37)
		// Understand this block and use it around the site to access the user id
		// with the ID, you can use the context object of golang to pass around any user details you need in any function
		// READ ABOUT using context with go

		// Retrieve the JWT token from cookies
		cookie, err := r.Cookie("authToken")
		if err != nil || cookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Parse and validate the JWT token
		tokenString := cookie.Value
		user, err := auth.ParseToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Printf("Authenticated user: %+v\n", user.Name)
		fmt.Printf("Authenticated user: %+v\n", user.ID)

		/* ------------------ TO HERE -----------------------------*/


		// Set the content type for the response
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Serve the dashboard page if the user is authenticated

		if err := tmpl.ExecuteTemplate(w, "dashboard.html", user.Name); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
