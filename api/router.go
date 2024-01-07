package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// initialize router
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// /all
	router.HandleFunc("/all", AllHandler).Methods("GET")
	router.HandleFunc("/random", RandomHandler("all")).Methods("GET")

	// /promote_communication
	router.HandleFunc("/promote_communication/random", RandomHandler("p_c")).Methods("GET")

	// /build_relationship
	router.HandleFunc("/build_relationship/random", RandomHandler("b_r")).Methods("GET")

	return router
}
