package api

import (
	"fmt"
	"net/http"

	auth "diawise/internal/auth"

	"gorm.io/gorm"
)

func Index(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//auth.RegisterUser(db, "toni", "toni@mail.com", "antony102")
		auth.LoginUser(db, "toni", "antony102")
		fmt.Fprintf(w, "Hello")
	}
}
