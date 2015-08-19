package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// START OMIT
func WithDB(s *mgo.Session, h http.Handler) http.Handler {
	return &dbwrapper{dbSession: s, h: h}
}

type dbwrapper struct {
	dbSession *mgo.Session
	h         http.Handler
}

func (h *dbwrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// copy the session
	dbcopy := h.dbSession.Copy()
	defer dbcopy.Close() // clean up after

	// put the session in the context for this Request
	context.Set(r, "db", dbcopy)

	// serve the request
	h.h.ServeHTTP(w, r)
}

// END OMIT

func main() {
	db, _ := mgo.Dial("localhost")
	defer db.Close()
	router := mux.NewRouter()
	router.Handle("/things", WithDB(db, http.HandlerFunc(handleThingsRead)))
	http.ListenAndServe(":8080", router)
}

func handleThingsRead(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)
	var result []interface{}
	db.DB("database").C("things").Find(nil).All(&result)
}
