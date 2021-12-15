package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sedessapi/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Mary Immaculate Conception")
}

// func getAllCustomers(w http.ResponseWriter, r *http.Request) {
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	//customers := []Customer{
	//	{Name: "Javier", City: "Madrid", Zipcode: "136"},
	//	{Name: "Alvaro", City: "Toledo", Zipcode: "963"},
	//}

	w.Header().Add(
		"Content-Type",
		"application/json",
	)
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
