package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	Method       string
	Path         string
	Body         io.Reader
	BodyContains string
	Status       int
}{{
	Method:       "GET",
	Path:         "/things",
	BodyContains: "Hello Golang UK Conference",
	Status:       http.StatusOK,
}, {
	Method:       "POST",
	Path:         "/things",
	Body:         strings.NewReader(`{"name":"Golang UK Conference"}`),
	BodyContains: "Hello Golang UK Conference",
	Status:       http.StatusCreated,
}}

// START OMIT
func TestAll(t *testing.T) {
	assert := assert.New(t)
	server := httptest.NewServer(&myhandler{}) // HL
	defer server.Close()                       // HL
	for _, test := range tests {
		r, err := http.NewRequest(test.Method, server.URL+test.Path, test.Body) // HL
		assert.NoError(err)
		// call handler
		response, err := http.DefaultClient.Do(r) // HL
		assert.NoError(err)
		actualBody, err := ioutil.ReadAll(response.Body)
		assert.NoError(err)
		// make assertions
		assert.Contains(actualBody, test.BodyContains, "body")        // HL
		assert.Equal(test.Status, response.StatusCode, "status code") // HL
	}
}

// END OMIT
type myhandler struct{}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestAll", F: TestAll})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
