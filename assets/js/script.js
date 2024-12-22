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
        status: "Doing",
        priority: priority
    };

    // Send the data to the API using the fetch API
    fetch('http://localhost:8080/api/addTask', {
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