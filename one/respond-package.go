// START OMIT
import (
	"net/http"

	"github.com/matryer/respond"
)

func handleSomething(w http.ResponseWriter, r *http.Request) {
	data, err := loadFromDB()
	if err != nil {
		// respond with an error
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}
	// respond with OK, and the data
	respond.With(w, r, http.StatusOK, data)
}

// END OMIT
