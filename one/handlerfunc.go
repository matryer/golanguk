package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type msg struct {
	Message string `json:"msg"`
}

// START OMIT
func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/goodbye", handleBye)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&msg{Message: "Hello Golang UK Conf"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleBye(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&msg{Message: "See you next year"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// END OMIT
