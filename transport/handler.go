package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcusziade/planets-rest-api/planet"
)

type Handler struct {
	Router  *mux.Router
	Service *planet.Service
}

type Response struct {
	Message string
	Error   string
}

// Returns a pointer to a `Handler`
func NewHandler(s *planet.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// Setup routes for the REST API
func (h *Handler) CreateRoutes() {
	fmt.Println("Creating routes")
	h.Router = mux.NewRouter()

}

func sendErrorResponse(w http.ResponseWriter, message string, error error) {
	w.WriteHeader(http.StatusInternalServerError)
	r := Response{Message: message, Error: error.Error()}
	if error := json.NewEncoder(w).Encode(r); error != nil {
		panic(error)
	}
}
