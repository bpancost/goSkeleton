package rest

import (
	"encoding/json"
	"net/http"

	"goSkeleton/internal/logging"
)

type AddPersonRequest struct {
	Name string `json:"name"`
}

type AddPersonResponse struct {
	ID string `json:"id"`
}

func (adapter Adapter) AddPerson(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetRequestLogger(req)
	var jsonRequest AddPersonRequest
	err := json.NewDecoder(req.Body).Decode(&jsonRequest)
	if err != nil {
		logging.Errorf("failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := adapter.Usecases.AddPersonCase(jsonRequest.Name)
	if err != nil {
		logger.Errorf("failed to add person: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := AddPersonResponse{ID: id}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Errorf("failed to encode response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
