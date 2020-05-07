package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"webServerGo/entity"

	"github.com/gorilla/mux"
)

func UpdateCount(w http.ResponseWriter, r *http.Request) {

	cartId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	var updatedCart entity.ShoppingBasket

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter in the appropriate count information")
	}
	json.Unmarshal(reqBody, &updatedCart)

	for i, cart := range entity.Basket {
		if cart.ID == cartId {
			cart.Count = updatedCart.Count
			entity.Basket = append(entity.Basket[:i], cart)
			json.NewEncoder(w).Encode(cart)
		}
	}
}
