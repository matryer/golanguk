package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT
func main() {

	router := mux.NewRouter()

	// adapt a single route
	router.Handle("/", Adapt(myHandler, WithHeader("X-Something", "Specific")))

	// adapt all handlers
	http.Handle("/", Adapt(router,
		SupportXHTTPMethodOverride(),
		WithHeader("Server", "MyApp v1"),
		Logging(logger),
	))
}

// END OMIT
