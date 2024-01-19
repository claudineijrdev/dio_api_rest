package api

import (
	"dio_api_rest/src/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customer.ID = fmt.Sprint(len(model.Customers) + 1)
	model.Customers = append(model.Customers, customer)
	w.WriteHeader(http.StatusCreated)
}

func ListCustomersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Customers)
}

func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, customer := range model.Customers {
		if customer.ID == id {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var uptCustomer model.Customer

	err := json.NewDecoder(r.Body).Decode(&uptCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, customer := range model.Customers {
		if customer.ID == id {
			uptCustomer.ID = id
			model.Customers[i] = uptCustomer
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, customer := range model.Customers {
		if customer.ID == id {
			model.Customers = append(model.Customers[:i], model.Customers[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
