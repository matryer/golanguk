package main

import (
	"log"
	"net/http"

	"github.com/matryer/respond"
)

// START OMIT
func main() {

	opts := &respond.Options{
		OnErr: func(err error) {
			log.Println("something went wrong:", err)
		},
		Before: func(w http.ResponseWriter, r *http.Request, status int, data interface{}) (int, interface{}) {
			w.Header().Set("Server", "My App Server")
			if err, ok := data.(error); ok {
				return status, map[string]interface{}{"error": err.Error()}
			}
			return status, data
		},
		After: func(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
			log.Println("replied with", status, data)
		},
	}

	// wrap existing handler with opts.Handler
	http.Handle("/foo", opts.Handler(fooHandler)) // HL
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT

var fooHandler http.Handler
