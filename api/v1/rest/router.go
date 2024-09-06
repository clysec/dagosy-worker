package rest

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/request", MakeRequest).Methods("GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS")
}
