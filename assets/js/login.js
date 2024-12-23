const form = document.getElementById("login-form");

// Add an event listener to handle form submission
form.addEventListener("submit", function (event) {
  event.preventDefault(); // Prevent the default form submission behavior

  // Get the username and password from the form
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  // Prepare the data for the POST request
  const data = {
    username: username,
    password: password,
  };

  // Send the POST request
  fetch("http://localhost:8080/api/user/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json()) // Parse the JSON response
    .then((data) => {
      if (data.token) {
        // Handle successful login (assuming the response contains the token)
        alert("Login successful! Redirecting...");
        
        // Add Token to cookie
        document.cookie = `jwt_token=${data.token}; path=/; Secure; SameSite=Strict`;

        // Redirect to the main page (http://localhost:8080) with the token in the headers
        //window.location.replace("http://localhost:8080");
        window.open("http://localhost:8080", "_self");
      } else {
        // Handle error (assuming error response contains the 'Message')
        alert("Login failed: " + data.Message);
      }
    })
    .catch((error) => {
      // Handle any network or other errors
      console.error("Error during login:", error);
      alert("An error occurred while processing your login request.");
    });
});


