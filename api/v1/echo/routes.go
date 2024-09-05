package echo

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", UniversalEcho).Methods("GET", "PUT", "PATCH", "POST", "DELETE")
}
