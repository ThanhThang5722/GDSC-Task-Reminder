const form = document.getElementById("signup-form");
var link2 = "https://gdsc-task-reminder.onrender.com"
// Add an event listener to handle form submission
form.addEventListener("submit", function (event) {
  event.preventDefault(); // Prevent the default form submission behavior

  // Get the username and password from the form
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;
  const email = document.getElementById("email").value;

  // Prepare the data for the POST request
  const data = {
    username: username,
    password: password,
    email: email
  };
  var connect = link2 + "/api/user/signup"
  // Send the POST request
  fetch(connect, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json()) // Parse the JSON response
    .then(() => {
        alert("Signup successful! Redirecting...");
        var connect = link2 + "/login"
        window.location.href = connect;
    })
    .catch((error) => {
      // Handle any network or other errors
      console.error("Error during SignUp:", error);
      alert("An error occurred while processing your Sign Up request.");
    });
});
