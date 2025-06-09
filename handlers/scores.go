package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/bakageddy/grades/data"
)

type HomeHandlers struct {
	DB data.Database
}

func (h *HomeHandlers) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/score/add", h.AddEntry)
	mux.HandleFunc("/score/delete", h.DeleteEntry)
	mux.HandleFunc("/score/update", h.UpdateEntry)
	mux.HandleFunc("/score/calculate", h.CalculateScore)
	
	return mux
}

func (h *HomeHandlers) AddEntry(w http.ResponseWriter, r *http.Request) {
	register_no := r.FormValue("register_no")
	if register_no == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Failed to obtain register_no")
	}

	grade := r.FormValue("grade")
	if grade == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Failed to obtain grade")
	}

	credits := r.FormValue("credits")
	if credits == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Failed to obtain credits")
	}

	subject_code := r.FormValue("subject_code")
	if subject_code == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Failed to obtain subject_code")
	}

	io.WriteString(w, "TODO")
}

func (h *HomeHandlers) DeleteEntry(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO")
}

func (h *HomeHandlers) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO")
}

func (h *HomeHandlers) CalculateScore(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO")
}
