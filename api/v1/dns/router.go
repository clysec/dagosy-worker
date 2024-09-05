package dns

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/request/{type}/{domain}", DNSRequest).Methods("GET")
	router.HandleFunc("/reverse/{ip}", ReverseDNSRequest).Methods("GET")
}
