package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// START OMIT
func TestHandler(t *testing.T) {
	assert := assert.New(t) // testify

	// make test w and r
	r, err := http.NewRequest("GET", "/", nil)
	assert.NoError(err)
	w := httptest.NewRecorder()

	// call handler
	handleSomething(w, r)

	// make assertions
	assert.Equal(w.Body.String(), "Hello Golang UK Conference", "body")
	assert.Equal(w.Code, http.StatusOK, "status code")
}

func handleSomething(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

// END OMIT

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestHandler", F: TestHandler})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
