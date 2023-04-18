package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.SearchUser).Methods(http.MethodGet)

	fmt.Println("Listen on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
