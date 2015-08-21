package main

import (
	"log"

	"net/http"
)

// START OMIT
func WithLogging(l *log.Logger, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println(r.Method, r.URL.Path)
		h(w, r)
	}
}

// END OMIT

func main() {}
