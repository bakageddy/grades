package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bakageddy/grades/handlers"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	file, err := os.Open("schema.sql")
	if err != nil {
		log.Fatalf("Failed to process migrations: %s\n", err.Error())
		return
	}

	body, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to process migrations: %s\n", err.Error())
		return
	}

	database, err := sql.Open("sqlite", "file.db")
	if err != nil {
		log.Fatalf("Failed to open database: %s\n", err.Error())
		return
	}

	_, db_err := database.Exec(string(body))
	if db_err != nil {
		log.Fatalf("Failed to apply migrations: %s\n", db_err.Error())
		return
	}

	auth_handle := handlers.Handlers {
		DB: database,
	}

	top_level_mux := http.NewServeMux()
	// This trailing slash is important i guess
	top_level_mux.Handle("/auth/",
		http.StripPrefix("/auth", auth_handle.GetMux()),
	)

	top_level_mux.HandleFunc("/ping", handlers.Ping)

	server_err := http.ListenAndServe("127.0.0.1:42069", top_level_mux);
	if (server_err != nil) {
		log.Fatalln(server_err.Error())
	}
	return;
}
