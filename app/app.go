package app

import (
	"github.com/gorilla/mux"
	"github.com/sedessapi/banking/domain"
	"github.com/sedessapi/banking/service"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:7000", router))
}
