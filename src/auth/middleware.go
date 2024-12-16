package auth

import (
	"net/http"
    "context"
    "fmt"
    "strings"
    "github.com/dgrijalva/jwt-go"
)

type key string
const userKey key = "user"


var mySigningKey = []byte("secret")

// Function to parse and validate the JWT
func ParseToken(tokenString string) (User, error) {
	// fmt.Println("========================================")
	// fmt.Println("ONE")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Ensure token signing method is expected
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return mySigningKey, nil
    })
	// fmt.Println("TWO: ", token)


    if err != nil {
        return User{}, err
    }
	// fmt.Println("THREE: ", token)


    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return User{}, fmt.Errorf("invalid token")
    }
	// fmt.Println("FOUR: ", token)


    // Safely parse the claims
    userID, ok := claims["id"].(float64)
    if !ok {
        return User{}, fmt.Errorf("missing or invalid user ID")
    }
	// fmt.Println("FIVE id => : ", userID)


    user := User{
        ID:   uint(userID), // Convert float64 to uint
        Name: claims["name"].(string),
        Email: claims["email"].(string),
    }
	// fmt.Println("SIX: ", claims)

	// fmt.Println("========================================")
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


func helloHandler(w http.ResponseWriter, r *http.Request) {
    user, ok := r.Context().Value(userKey).(User)
    if !ok {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "Hello, %s!", user.Name)
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
