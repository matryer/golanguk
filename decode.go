package main

import (
	"encoding/json"

	"github.com/matryer/respond"

	"net/http"
)

// START OMIT
type gopher struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func handleCreateGopher(w http.ResponseWriter, r *http.Request) error {
	var g gopher
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		return err
	}
	if err := SaveGopher(&g); err != nil {
		return err
	}
	respond.With(w, r, http.StatusOK, &g)
	return nil
}

// END OMIT

func main() {}

func SaveGopher(gopher interface{}) error {
	return nil
}
