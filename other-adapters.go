package main

import "net/http"

// START OMIT

// WithHeader is an Adapter that sets an HTTP handler.
func WithHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header.Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}

// SupportXHTTPMethodOverride adds support for the X-HTTP-Method-Override
// header.
func SupportXHTTPMethodOverride() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := r.Header.Get("X-HTTP-Method-Override")
			if len(m) > 0 {
				r.Method = m
			}
			h.ServeHTTP(w, r)
		})
	}
}

// END OMIT
