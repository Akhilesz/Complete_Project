package main

import (
	"database/sql"
	"log"
	"net/http"
	"url/internal"
	"url/repository"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := &repository.URLRepo{DB: db}
	h := &internal.Handler{Repo: repo}

	http.HandleFunc("/shorten", h.ShortenUrl)
	http.HandleFunc("/", h.Redirect)

	log.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
