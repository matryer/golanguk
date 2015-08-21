package main

import "net/http"

// START OMIT

func handleThingsRead(w http.ResponseWriter, r *http.Request, logger *Logger) { /*...*/ }

func main() {
	http.HandleFunc("/", handleThingsRead)
}

// END OMIT

type Logger struct{}
