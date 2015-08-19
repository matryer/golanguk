package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START OMIT

func main() {
	http.HandleFunc("/", handleError(handleSomethingWithErr))
}

func handleError(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println("http error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleSomethingWithErr(w http.ResponseWriter, r *http.Request) error {
	thing, err := GetThing()
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(thing)
}

// END OMIT
func GetThing() (interface{}, error) {
	return nil, nil
}
