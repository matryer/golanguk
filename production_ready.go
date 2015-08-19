package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START OMIT
type msg struct {
	Message string `json:"msg"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(&msg{Message: "Hello Golang UK Conf"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

// END OMIT
