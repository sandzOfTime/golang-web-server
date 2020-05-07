package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webServerGo/entity"

	"github.com/gorilla/mux"
)

func GetOnePublisher(w http.ResponseWriter, r *http.Request) {

	publisherId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	for _, pub := range entity.Publishers {

		if pub.ID == publisherId {
			json.NewEncoder(w).Encode(pub)
		}
	}

}

func GetAllPublishers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entity.Publishers)
}
