package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateCount(w http.ResponseWriter, r *http.Request) {

	cartId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

}
