package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}
