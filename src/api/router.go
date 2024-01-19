package api

import "github.com/gorilla/mux"

func ConfigServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/customers", AddCustomerHandler).Methods("POST")
	r.HandleFunc("/customers", ListCustomersHandler).Methods("GET")
	r.HandleFunc("/customers/{id}", GetCustomerHandler).Methods("GET")
	r.HandleFunc("/customers/{id}", UpdateCustomerHandler).Methods("PUT")
	r.HandleFunc("/customers/{id}", DeleteCustomerHandler).Methods("DELETE")
	return r
}
