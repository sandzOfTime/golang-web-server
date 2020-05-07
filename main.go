package main

import (
	"fmt"
	"log"
	"net/http"
	"webServerGo/handler"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/customers/{id}", handler.GetOneCustomer).Methods("GET")
	router.HandleFunc("/customers", handler.GetCustomers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))

}
