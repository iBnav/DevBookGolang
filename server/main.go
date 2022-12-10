package main

import (
	"fmt"
	"log"
	controller "modulo/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.PostUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
