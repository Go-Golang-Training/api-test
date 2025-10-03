package main

import (
	"log"
	"net/http"

	"github.com/Go-Golang-Training/toolkit/v2"
)

func main() {
	mux := routes()
	log.Println("Starting server (Application) on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/api/login", login)
	mux.HandleFunc("/api/logout", logout)

	return mux
}

func login(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Login Endpoint"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "You have been logged out",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
