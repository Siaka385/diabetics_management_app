package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"diawise/internal/api"
	"diawise/internal/repository"
	"diawise/internal/shared"
)

var (
	db           *gorm.DB
	sessionStore *sessions.CookieStore
)

func init() {
	db = repository.InitializeDatabase("data/diawise.db")
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
}

func main() {
	port := shared.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := api.Router(db)

	http.ListenAndServe(portStr, router)
}
