package main

import (
	"net/http"
)

// START OMIT
func Wrap(h http.Handler) http.Handler {
	return &wrapper{handler: h}
}

type wrapper struct {
	handler http.Handler
}

func (h *wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: do something before each request

	h.handler.ServeHTTP(w, r)

	// TODO: do something after each request
}

func main() {
	handler := NewMyHandler()
	http.Handle("/", Wrap(handler))
}

// END OMIT
