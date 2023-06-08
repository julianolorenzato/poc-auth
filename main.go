package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/julianolorenzato/poc-auth/database"
	"github.com/julianolorenzato/poc-auth/handlers"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var db *sql.DB

func init() {
	var err error

	db, err = database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialised")
}

func main() {
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello friend"))
	})

	r.Post("/register", handlers.RegisterUser(db))
	r.Post("/authenticate", handlers.AuthenticateUser(db))

	log.Println("Starting server...")

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Stoped serving new connections")

	http.ListenAndServe(":8080", r)
}
