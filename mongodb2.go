package main

import (
	"net/http"

	"github.com/matryer/respond"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func WithDB(s *mgo.Session, h http.Handler) http.Handler {
	return &dbwrapper{dbSession: s, h: h}
}

type dbwrapper struct {
	dbSession *mgo.Session
	h         http.Handler
}

func (dbwrapper *dbwrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// copy the session
	dbcopy := h.dbSession.Copy()
	defer dbcopy.Close() // clean up after

	// put the session in the context for this Request
	context.Set(r, "db", dbcopy)

	// serve the request
	dbwrapper.h.ServeHTTP(w, r)
}

// START OMIT
func main() {
	// connect to mongo
	db, _ := mgo.Dial("localhost") // expensive
	defer db.Close()

	router := mux.NewRouter()
	router.Handle("/things", WithDB(db, http.HandlerFunc(handleThingsRead))) // HL
	router.Handle("/status", http.HandlerFunc(handleStatus))

	http.ListenAndServe(":8080", context.ClearHandler(router)) // HL
}

func handleThingsRead(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session) // HL

	var results []interface{}
	if err := db.DB("myapp").C("things").Find(nil).All(&results); err != nil {
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.With(w, r, http.StatusOK, results)
}

// END OMIT

func handleStatus(w http.ResponseWriter, r *http.Request) {}
