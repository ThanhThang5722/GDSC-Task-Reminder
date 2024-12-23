var link1 = "https://gdsc-task-reminder.onrender.com"
// Get the modal
var modal = document.getElementById('formModal');

// Get the "Add" button that opens the modal
var addButton = document.getElementById('add-button');

// Get the close button (X)
var closeButton = document.getElementById('closeButton');

// Get the form
var form = document.getElementById('taskForm');

// When the user clicks the "Add Task" button, open the modal
addButton.onclick = function() {
    modal.style.display = "block"; // Show the modal
}

// When the user clicks the "X" button, close the modal
closeButton.onclick = function() {
    modal.style.display = "none"; // Hide the modal
}

// When the user clicks anywhere outside the modal, close it
window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none"; // Hide the modal
    }
}

// Handle the form submission
form.onsubmit = function(event) {
    event.preventDefault(); // Prevent form submission (for demonstration)

    // Get the values from the form
    var title = document.getElementById('title').value;
    var description = document.getElementById('description').value;
    var deadline = document.getElementById('deadline').value;
    var priority = document.getElementById('priority').value;
    deadline = new Date(deadline).toISOString();
    // Create the data object
    var taskData = {
        title: title,
        description: description,
        deadline: deadline,
        priority: priority
    };
    var connect = link1 + "/api/Task"
    // Send the data to the API using the fetch API
    fetch(connect, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(taskData) // Convert the taskData object to JSON
    })
    .then(response => response.json()) // Parse the JSON response from the API
    .then(data => {
        console.log('Task added successfully:', data); // Handle success response
        modal.style.display = "none"; // Hide the modal after submission
        form.reset(); // Reset the form fields

        var taskCard = createTaskCard(data.task);
        console.log(data)
        // Append task to the appropriate quadrant based on priority
        appendTaskToQuadrant(taskCard, taskData.priority);
    })
    .catch(error => {
        console.error('Error adding task:', error); // Handle error response
    });
};

function createTaskCard(task) {
    var card = document.createElement('div');
    card.classList.add('card');
    card.innerHTML = `
        <h4>${task.title}</h4>
        <p>${task.description}</p>
        <small>Deadline: ${new Date(task.deadline).toLocaleString()}</small>
        <button class="done-btn" onclick="markAsDone('${task.title}', '${task.description}', '${task.deadline}')">
                            Done
                        </button>
    `;
    return card;
}

function appendTaskToQuadrant(card, priority) {
    var quadrant;

    // Assign quadrant based on the task priority
    switch (priority) {
        case "DoFirst":
            quadrant = document.getElementById('top-left');
            break;
        case "DoLater":
            quadrant = document.getElementById('top-right');
            break;
        case "Delegate":
            quadrant = document.getElementById('bottom-left');
            break;
        case "Eliminate":
            quadrant = document.getElementById('bottom-right');
            break;
        default:
            quadrant = document.getElementById('top-left'); // Default to top-left
    }

    // Append the card to the selected quadrant
    quadrant.appendChild(card);
}

document.getElementById("body").addEventListener("click", function(event) {
    // Check if the clicked element is a "Done" button
    if (event.target && event.target.classList.contains("done-btn")) {
        // Call markAsDone with the necessary details
        const button = event.target;
        const card = button.closest('.card');
        
        const title = card.querySelector('h4').textContent;
        const description = card.querySelector('p').textContent;
        const deadline = card.querySelector('small').textContent.replace('Deadline: ', '');
        
        markAsDone(button, title, description, deadline);
    }
});

function markAsDone(button, title, description, deadline) {
    const data = {
        title: title,
        description: description
    };
    var connect = link1 + "/api/Task"
    // Make the API call to mark the task as done
    fetch(connect, {
        method: "DELETE",  // or "PATCH"
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
    .then((response) => {
        if (response.ok) {
            alert("Task marked as done");

            // Find the closest parent with the 'card' class and remove it
            const card = button.closest('.card');
            if (card) {
                card.remove();  // Remove the task card from the UI
            }
        } else {
            alert("Error marking task as done");
        }
    })
    .catch((error) => {
        console.error("Error:", error);
    });
}

// Function to handle sign-out
function signOut() {
    // Remove the jwt_token cookie
    document.cookie = "jwt_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    
    // Redirect to login page
    var connect = link1 + "/login"
    window.location.href = connect;
}

// Attach event listener to the signout button
document.getElementById('signoutBtn').addEventListener('click', signOut);
