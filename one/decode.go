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

func handleCreateGopher(w http.ResponseWriter, r *http.Request) {
	var g gopher
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		respond.With(w, r, http.StatusBadRequest, err)
		return
	}
	if err := SaveGopher(&g); err != nil {
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}
	respond.With(w, r, http.StatusOK, &g)
}

// END OMIT

func main() {}

func SaveGopher(gopher interface{}) error {
	return nil
}
