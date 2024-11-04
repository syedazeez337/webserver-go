package server

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
