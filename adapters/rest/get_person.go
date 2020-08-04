package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"goSkeleton/internal/logging"
)

func (adapter Adapter) GetPerson(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetRestRequestLogger(req)
	vars := mux.Vars(req)
	id := vars["id"]
	person, err := adapter.Usecases.GetPersonCase(id)
	if err != nil {
		logger.Errorf("failed to get person: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		logger.Errorf("failed to encode response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
