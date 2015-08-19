package main

import (
	"net/http"

	"gopkg.in/mgo.v2"
)

// START OMIT

func handleThingsRead(w http.ResponseWriter, r *http.Request, db *mgo.Session) { /*...*/ }

func main() {
	http.HandleFunc("/", handleThingsRead)
}

// END OMIT
