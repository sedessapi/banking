package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Maria")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Javier", City: "Madrid", Zipcode: "136"},
		{Name: "Alvaro", City: "Valencia", Zipcode: "963"},
	}
	w.Header().Add(
		"Content-Type",
		"application/json",
	)
	json.NewEncoder(w).Encode(customers)
}
