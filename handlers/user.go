package handlers

type User struct {
	ID       string
	Username string
	Email    string
	Password string `json:"-"`
}
