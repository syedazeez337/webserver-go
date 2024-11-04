package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `
        <html>
            <head>
                <title>Welcome to Go Web Server</title>
                <style>
                    body { font-family: Arial, sans-serif; margin: 40px; }
                    h1 { color: #333; }
                </style>
            </head>
            <body>
                <h1>Welcome to Go Web Server</h1>
                <p>This is a simple web server built with Go!</p>
                <ul>
                    <li><a href="/about">About</a></li>
                    <li><a href="/users">Users</a></li>
                    <li><a href="/api/status">API Status</a></li>
                </ul>
            </body>
        </html>
    `)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
        <html>
            <head>
                <title>About - Go Web Server</title>
                <style>
                    body { font-family: Arial, sans-serif; margin: 40px; }
                    h1 { color: #333; }
                </style>
            </head>
            <body>
                <h1>About</h1>
                <p>This is a tutorial web server built with Go.</p>
                <p><a href="/">Back to Home</a></p>
            </body>
        </html>
    `)
}


func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User {
		{FirstName: "John", LastName: "Baker", Email: "john@gmail.com"},
		{FirstName: "Mono", LastName: "Smith", Email: "mono@gmail.com"},
	}

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case "POST":
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Message{
			Status: "success",
			Message: "User created successfully",
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ApiStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := Message {
		Status: "operational",
		Message: "API is running normally",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}