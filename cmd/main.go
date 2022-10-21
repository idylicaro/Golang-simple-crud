package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Profile  *Profile `json:"profile"`
}

type Profile struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var accounts []Account

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func main() {
	r := mux.NewRouter()

	accounts = append(accounts, Account{ID: "1", Username: "admin", Email: "admin@admin", Password: "admin", Profile: &Profile{ID: "1", Firstname: "admin", Lastname: "admin"}})

	r.HandleFunc("/", getHealthCheck).Methods("GET")
	r.HandleFunc("/accounts", getAccounts).Methods("GET")
	// r.HandleFunc("/accounts/{id}", getAccount).Methods("GET")
	// r.HandleFunc("/accounts", createAccount).Methods("POST")
	// r.HandleFunc("/accounts/{id}", updateAccount).Methods("PUT")
	// r.HandleFunc("/accounts/{id}", deleteAccount).Methods("DELETE")

	fmt.Printf("Starting server on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
