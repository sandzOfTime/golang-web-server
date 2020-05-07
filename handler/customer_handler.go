package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webServerGo/entity"

	"github.com/gorilla/mux"
)

func GetOneCustomer(w http.ResponseWriter, r *http.Request) {

	customerIdString := mux.Vars(r)["id"]
	customerId, err := strconv.Atoi(customerIdString)
	if err != nil {
		log.Fatal(err)
	}

	for _, cus := range entity.Customers {

		if cus.ID == customerId {
			json.NewEncoder(w).Encode(cus)
		}
	}

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(entity.Customers)

}
