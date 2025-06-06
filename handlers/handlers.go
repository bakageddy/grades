package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/search", Search)
	mux.HandleFunc("/ping", Ping)
	mux.HandleFunc("/register", h.Register)
	mux.HandleFunc("/login", h.Login)
	mux.HandleFunc("/logout", h.Logout)

	return mux
}

func Search(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(n)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
