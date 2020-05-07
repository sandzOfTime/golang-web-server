package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webServerGo/entity"

	"github.com/gorilla/mux"
)

func GetOneAuthor(w http.ResponseWriter, r *http.Request) {

	authorIdString := mux.Vars(r)["id"]
	authorId, err := strconv.Atoi(authorIdString)
	if err != nil {
		log.Fatal(err)
	}

	for _, cus := range entity.Authors {

		if cus.ID == authorId {
			json.NewEncoder(w).Encode(cus)
		}
	}

}

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(entity.Authors)

}
