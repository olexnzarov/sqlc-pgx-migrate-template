package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/olexnzarov/sqlc-pgx-migrate-template/internal/db"
	"github.com/olexnzarov/sqlc-pgx-migrate-template/internal/db/repositories/authors"
)

func main() {
	db, err := db.Setup(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to initialize the database: %s", err)
	}
	defer db.Close()

	server := newServer(db)
	server.listen()
}

type server struct {
	http    *http.Server
	authors *authors.Queries
}

func newServer(db *db.Database) *server {
	mux := http.NewServeMux()
	server := &server{
		http: &http.Server{
			Addr:    os.Getenv("LISTEN_ADDRESS"),
			Handler: mux,
		},
		authors: authors.New(db),
	}
	mux.HandleFunc("/authors", server.handleAuthors)
	return server
}

func (s *server) listen() {
	log.Printf("listening on %s", s.http.Addr)
	err := s.http.ListenAndServe()
	log.Printf("server stopped: %s", err)
}

func (s *server) handleAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := s.authors.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error 'GET /authors': %s", err)
		return
	}

	response, _ := json.Marshal(authors)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(response)
}
