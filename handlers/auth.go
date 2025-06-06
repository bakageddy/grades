package handlers

import (
	"log"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	register_no := r.FormValue("register_no")

	if register_no == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Register -> FormValue : Empty Register Number")
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Register -> DB.Begin : %s\n", err.Error())
		return
	}

	stmt, err := tx.Prepare("INSERT INTO Users(register_no) VALUES(?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Register -> tx.Prepare : %s\n", err.Error())
		return
	}

	result, err := stmt.Exec(register_no)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Register -> stmt.Exec : %s\n", err.Error())
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Register -> result.RowsAffected : %s\n", err.Error())
		return
	}

	if n != 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Register -> result.RowsAffected != 1")
		return
	} else {
		log.Printf("INFO: User %s Registered!\n", register_no)
		_ = tx.Commit()
		return
	}
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	register_no := r.FormValue("register_no")
	if register_no == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Panicln("ERR: Login -> register_no param is empty")
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Login -> DB.Begin : %s\n", err.Error())
		return
	}

	stmt, err := tx.Prepare("SELECT 1 FROM Users WHERE register_no = ?;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Login -> tx.Prepare : %s\n", err.Error())
		return
	}

	result := stmt.QueryRow(register_no)
	var student_register_count int
	if err := result.Scan(&student_register_count); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERR: Login -> Scan : %s\n", err.Error())
		return
	}

	if student_register_count != 1 {
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		log.Println("ERR: Login -> Count : What am i doing?")
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("INFO: Login -> Successful")
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
