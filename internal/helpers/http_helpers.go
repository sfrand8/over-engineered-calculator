package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse is used to define error response structure
// @Description Error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, errorString string, responseCode int) {
	w.WriteHeader(responseCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Error: errorString})
	if err != nil {
		log.Printf("Error when encoding error response: %s", err)
		http.Error(w, "something went wrong internally", http.StatusInternalServerError)
		return
	}
}
