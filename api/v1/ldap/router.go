package ldap

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/query", RunQuery).Methods("POST")
}
