window.onload = () => {

  /** Sign up */
  async function sendSignUpData() {
      let signupForm = document.getElementById("signup-form");
      const formData = new FormData(signupForm);
    
      // Convert FormData to a plain object
      const formValues = Object.fromEntries(formData.entries());
    
      try {
          const response = await fetch("http://localhost:8089/", {
              method: "POST",
              headers: {
                  "Content-Type": "application/json",
              },
              // Convert the object to a JSON string
              body: JSON.stringify(formValues),
          });
    
          // Return the parsed JSON from the response
          return await response.json();
      } catch (error) {
          console.error(error);
          throw error;
      }
  }
  
  let signUpButton = document.getElementById("signup-button");
  signUpButton.addEventListener("submit", async (event) => {
      event.preventDefault();
    
      let passwordField = document.getElementById("password");
      let confirmPasswordField = document.getElementById("confirm-password");
      let warningMessage = document.getElementById("email-exists-alert");
      let warningTitle = document.getElementById("warning-message-title");
    
      // Compare the values of the password and confirm password fields
      if (passwordField.value === confirmPasswordField.value) {
          warningMessage.classList.add("hide"); // Hide the warning if previously shown
    
          try {
              const result = await sendSignUpData();
              console.log("Result from server:", result);
    
              // Use the result as needed
              if (result.success) {
                  alert("Form submitted successfully!");
              } else {
                  warningMessage.classList.remove("hide");
                  warningTitle.textContent = result.message || "Submission failed.";
              }
          } catch (error) {
              alert("Unable to submit the form. Please try again later.");
          }
      } else {
          // Show feedback when passwords do not match
          warningMessage.classList.remove("hide");
          warningTitle.textContent = "Your passwords do not match.";
      }
  });
  
  
  /** Login */
  async function sendLoginData() {
      let loginForm = document.getElementById("login-form");
      const formData = new FormData(loginForm);
    
      // Convert FormData to a plain object
      const formValues = Object.fromEntries(formData.entries());
    
      try {
          const response = await fetch("http://localhost:8089/", {
              method: "POST",
              headers: {
                  "Content-Type": "application/json",
              },
              // Convert the object to a JSON string
              body: JSON.stringify(formValues),
          });
    
          // Return the parsed JSON from the response
          return await response.json();
      } catch (error) {
          console.error(error);
          throw error;
      }
  }
  
  let loginButton = document.getElementById("login-button");
  loginButton.addEventListener("submit", async (event) => {
      event.preventDefault();
    
      let warningMessage = document.getElementById("login-error-alert");
    
      try {
          const result = await sendLoginData();
          console.log("Result from server:", result);
    
          // Use the result as needed
          if (result.success) {
              alert("Login successful!");
          } else {
              warningMessage.classList.remove("hide");
          }
      } catch (error) {
          alert("Unable to submit the form. Please try again later.");
      }
  });
  
  };
  