package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// UserContextKey is a key type for storing user in context
type UserContextKey string
type key string
const userKey key = "user"
var mySigningKey = []byte("secret")

// AuthMiddleware handles authentication and adds user to request context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Retrieve the JWT token from cookies
        cookie, err := r.Cookie("authToken")
        if err != nil || cookie == nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Parse and validate the JWT token
        tokenString := cookie.Value
        user, err := ParseToken(tokenString)
        if err != nil {
            fmt.Printf("Token parsing error: %v\n", err)
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Debug print
        // fmt.Printf("Middleware User: %+v\n", user)

        // Create a new context with the user information
        ctx := context.WithValue(r.Context(), userKey, user)
        
        // Create a new request with the updated context
        r = r.WithContext(ctx)

        // Call the next handler with the new request
        next.ServeHTTP(w, r)
    }
}

func GetUserFromContext(r *http.Request) (*User, bool) {
    // Retrieve the user from the context
    userValue := r.Context().Value(userKey)
    
    // Debug print
    // fmt.Printf("Context Value: %+v\n", userValue)

    // Type assert to *User
    user, ok := userValue.(*User)
    if !ok {
        fmt.Println("User not found in context or wrong type")
        return nil, false
    }

    return user, true
}

func ParseToken(tokenString string) (*User, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Ensure token signing method is expected
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return mySigningKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    // Safely parse the claims
    userID, ok := claims["id"].(float64)
    if !ok {
        return nil, fmt.Errorf("missing or invalid user ID")
    }

    user := &User{
        ID:    uint(userID), // Convert float64 to uint
        Name:  claims["name"].(string),
        Email: claims["email"].(string),
    }

    return user, nil
}

// Middleware to check JWT token and add user info to context
func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			// Redirect to login if no token is provided
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Remove 'Bearer ' prefix if it's there
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse and validate the token
		user, err := ParseToken(tokenString)
		if err != nil {
			// Redirect to login if token is invalid
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Store user information in the request context
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// https://github.com/Siaka385/diabetics_management_app
		next(w, r)
	}
}

// Middleware for handling POST requests with CORS and content type
func POST(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure it's a POST request
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

// Middleware for handling GET requests with CORS
func GET(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure it's a GET request
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
