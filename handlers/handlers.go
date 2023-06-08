package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterUserDTO struct {
	Username string
	Email    string
	Password string
}

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto RegisterUserDTO

		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		_, err = db.Exec(
			`INSERT INTO users (id, username, email, password)
			VALUES (?, ?, ?, ?)`, "asasasa", dto.Username, dto.Email, dto.Password)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

	}
}

type AuthenicateUserDTO struct {
	Username string
	Password string
}

func AuthenticateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto AuthenicateUserDTO

		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var user User

		row := db.QueryRow("SELECT * FROM users WHERE username = ?", dto.Username)
		fmt.Println(dto)

		err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if user.Password != dto.Password {
			http.Error(w, "invalid credentials", 500)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("OK")
	}
}

func ListUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		var users []User

		for rows.Next() {
			var user User

			err := rows.Scan(&user.ID, user.Username, user.Email)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			users = append(users, user)
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
