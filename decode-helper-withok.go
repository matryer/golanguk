package main

import (
	"encoding/json"
	"net/http"
)

// START OMIT
// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if validatable, ok := v.(ok); ok {
		return validatable.OK()
	}
	return nil
}

// END OMIT

type ok interface {
	OK() error
}
