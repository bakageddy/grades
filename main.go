package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bakageddy/grades/handlers"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	database, err := sql.Open("sqlite", "file.db")
	if err != nil {
		log.Fatalf("Failed to open database: %s\n", err.Error())
		return
	}

	// Test if the connection works!
	_ = database.QueryRow("SELECT 1 * 1;");

	auth_mux := handlers.GetMux();

	http.Handle("/", auth_mux)

	server_err := http.ListenAndServe("127.0.0.1:42069", nil);
	if (server_err != nil) {
		log.Fatalln(server_err.Error())
	}
	return;
}
