package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/glebarez/go-sqlite"
)

func search(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(n)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
}

func main() {
	database, err := sql.Open("sqlite", "file.db")
	if err != nil {
		log.Fatalf("Failed to open database: %s\n", err.Error())
		return
	}

	_ = database.QueryRow("SELECT 1 * 1;");

	http.HandleFunc("/auth", auth)
	http.HandleFunc("/search", search)
	http.HandleFunc("/ping", ping)

	server_err := http.ListenAndServe("127.0.0.1:42069", nil);
	if (server_err != nil) {
		log.Fatalln(server_err.Error())
	}
	return;
}
