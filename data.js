document.addEventListener("DOMContentLoaded", function () {
    const authForm = document.getElementById("authForm");

    authForm.addEventListener("submit", function (e) {
        e.preventDefault();

        const name = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        const userData = {
            name: document.getElementById("username").value,
            password: document.getElementById("password").value
        };

        fetch('http://localhost:8082/registerUser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        })
        .then(response => response.json())
        .then(data => {
            // Handle the response from the server
            console.log(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
})