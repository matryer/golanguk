package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/hello", handleHello)
	router.HandleFunc("/goodbye", handleGoodbye)

	router.HandleFunc("/things/{id}", handleThingsRead).Methods("GET")   // HL
	router.HandleFunc("/things/{id}", handleThingsUpdate).Methods("PUT") // HL

	http.Handle("/", router)
}

func handleThingsRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // HL
	// TODO: load thing with ID vars["id"]
}

// END OMIT
