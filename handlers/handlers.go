package handlers

import (
	"log"
	"net/http"
)

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

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO"))
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/search", Search)
	mux.HandleFunc("/ping", Ping)
	mux.HandleFunc("/auth", Auth)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)

	return mux
}
