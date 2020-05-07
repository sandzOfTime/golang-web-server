package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webServerGo/entity"

	"github.com/gorilla/mux"
)

func GetOneBook(w http.ResponseWriter, r *http.Request) {

	bookId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range entity.Books {

		if b.ID == bookId {
			json.NewEncoder(w).Encode(b)
		}
	}

}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(entity.Books)

}
