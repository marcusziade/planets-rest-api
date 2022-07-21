package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcusziade/planets-rest-api/planet"
)

// Get all planets from the database
func (h *Handler) GetAllPlanets(w http.ResponseWriter, r *http.Request) {
	planets, err := h.Service.GetAllPlanets()
	if err != nil {
		sendErrorResponse(w, "Failed to get planets", err)
	}

	if err := json.NewEncoder(w).Encode(planets); err != nil {
		panic(err)
	}
}

func (h *Handler) GetPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	planetID, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from planetID", err)
	}

	planet, err := h.Service.GetPlanet(uint(planetID))
	if err != nil {
		sendErrorResponse(w, "Error getting planet for ID", err)
	}

	if err := json.NewEncoder(w).Encode(planet); err != nil {
		panic(err)
	}
}

func (h *Handler) AddPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var planet planet.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body for new planet", err)
	}

	planet, err := h.Service.AddPlanet(planet)
	if err != nil {
		sendErrorResponse(w, "Failed to add new planet to database", err)
	}

	if err := json.NewEncoder(w).Encode(planet); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	planetID, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(w, "Failed to parse UINT for planetID", err)
	}

	var planet planet.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body for planet", err)
	}

	planet, err = h.Service.UpdatePlanet(uint(planetID), planet)
	if err != nil {
		sendErrorResponse(w, "Failed to update planet for ID", err)
	}

	if err := json.NewEncoder(w).Encode(planet); err != nil {
		panic(err)
	}
}

func (h *Handler) DeletePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	planetID, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(w, "Failed to parse UINT for planetID", err)
	}

	err = h.Service.DeletePlanet(uint(planetID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete planet for planetID", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Planet was deleted succesfully"}); err != nil {
		panic(err)
	}
}
