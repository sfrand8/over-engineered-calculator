package get_history

import (
	"encoding/json"
	http_middleware "github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	"log"
	"net/http"
	"over-engineered-calculator/internal/database"
	"over-engineered-calculator/internal/helpers"
)

type Response struct {
	CalculationHistory []HistoryEntry `json:"calculationHistory"`
}

type HistoryEntry struct {
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

// ErrorResponse is used to define error response structure
// @Description Error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Get calculation get_history for a user
// @Description This endpoint will retrieve the calculation get_history for a user based on the Authorization token
// @Tags calculations
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token (UUID)"
// @Success 200 {object} Response "Calculation get_history for the user"
// @Failure 500 {object} ErrorResponse "Error retrieving get_history"
// @Router /get_history [get]
func createHistoryHandler(calculationHistoryRetriever calculationHistoryRetriever) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := http_middleware.GetUserIDFromContext(r.Context())

		history, err := calculationHistoryRetriever.GetHistory(userID)
		if err != nil {
			log.Printf("Error when fetching Calculation History: %s", err)
			helpers.WriteErrorResponse(w, "something went wrong", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(mapResponse(history))
		if err != nil {
			log.Printf("Error when encoding history response: %s", err)
			helpers.WriteErrorResponse(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}

func mapResponse(calculations []database.Calculation) Response {
	var historyEntries []HistoryEntry
	for _, calculation := range calculations {
		historyEntries = append(historyEntries, HistoryEntry{
			Expression: calculation.Expression,
			Result:     calculation.Result,
		})
	}

	return Response{
		CalculationHistory: historyEntries,
	}
}
