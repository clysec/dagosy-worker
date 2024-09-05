package domain

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/whois/{domain}", Whois).Methods("GET")
	router.HandleFunc("/whois/{domain}/json", WhoisJSON).Methods("GET")

	router.HandleFunc("/nameservers/{domain}", GetNameservers).Methods("GET")
}
