package main

import (
	"net/http"
)

// START OMIT
func Wrap(h http.Handler) http.Handler { // HL
	return &wrapper{handler: h}
}

type wrapper struct {
	handler http.Handler
}

func (h *wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// TODO: do something before each request // HL
	h.handler.ServeHTTP(w, r) // HL
	// TODO: do something after each request // HL

}

func main() {
	handler := NewMyHandler()
	http.Handle("/", Wrap(handler)) // HL
}

// END OMIT
