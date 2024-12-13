window.onload = () => {
    async function sendData() {
        let form = document.getElementById("signup-form");
        const formData = new FormData(form);
      
        // Convert FormData to a plain object
        const data = Object.fromEntries(formData.entries());
      
        try {
          const response = await fetch("localhost:8089", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            // Convert the object to a JSON string
            body: JSON.stringify(data),
          });
      
          // Return the parsed JSON from the response
          return await response.json();
        } catch (e) {
          console.error(e);
          throw e;
        }
      }
      
  
    let signupform = document.getElementById("signup-form");
  
  signupform.addEventListener("submit", async (event) => {
      event.preventDefault();
  
      var password = document.getElementById("password");
      var confirmpassword = document.getElementById("confirm-password");
      var warnmsg = document.getElementById("email-exists-alert");
      var msgTitle = document.getElementById("WarningMessageTitle");
  
      // Compare the values of the password and confirm password fields
      if (password.value === confirmpassword.value) {
        warnmsg.classList.add("hide"); // Hide the warning if previously shown
  
        try {
          const result = await sendData();
          console.log("Result from server:", result);
  
          // Use the result as needed
          if (result.success) {
            alert("Form submitted successfully!");
          } else {
            warnmsg.classList.remove("hide");
          }
        } catch (e) {
          alert("Unable to submit the form. Please try again later.");
        }
      } else {
        // Show feedback when passwords do not match
        warnmsg.classList.remove("hide");
        msgTitle.textContent = "Your passwords do not match.";
        document.getElementById("warningBody").innerHTML=""
      }
    });
  };
  
  /**package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name            string `json:"name"`             // User's name
	Gmail           string `json:"gmail"`            // User's Gmail address
	Password        string `json:"password"`         // Password
	ConfirmPassword string `json:"confirm_password"` // Confirmation of the password
}

func test(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON into the User struct
	var user User
	if err := json.Unmarshal(content, &user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Process the user data (example: print it)
	fmt.Printf("Received user: %+v\n", user)

	// Example response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "message": "User data received successfully"}`))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", test)

	http.ListenAndServe(":8089", mux)
}
 */