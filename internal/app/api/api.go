package api

import (
	"database/sql"
	"hlservice-db/internal/app/storage/postgres"
	"log"
	"net/http"
)

func Start(config *Config) error {
	db, err := db(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	store := postgres.New(db)
	srv := New(store)

	return http.ListenAndServe(config.BindAddr, srv.router)
}

func db(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
