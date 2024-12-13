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
  
  