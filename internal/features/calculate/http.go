package calculate

import (
	"encoding/json"
	"errors"
	http_middleware "github.com/sfrand8/over-engineered-calculator/pkg/http-middleware"
	"log"
	"net/http"
	"over-engineered-calculator/internal/helpers"
)

// Request represents the structure of the request body for the calculation
// @Description Request body for calculation endpoint
// @Param expression body string true "Mathematical expression to evaluate"
type Request struct {
	Expression string `json:"expression"`
}

// Response represents the structure of the response for the calculation
// @Description Response body for calculation endpoint
type Response struct {
	// @Description The result of the mathematical expression evaluation
	// @Example "42.0" // Optional: Provide an example result
	Result string `json:"result,omitempty"`
}

// @Summary Calculate a simple expression
// @Description This endpoint will take a simple calculation string and return the result
// @Tags calculations
// @Accept json
// @Produce json
// @Param calculation body Request true "Calculation request body"
// @Param Authorization header string true "Authorization token (UUID)"
// @Success 200 {object} Response "Calculation result"
// @Failure 400 {object} helpers.ErrorResponse "Invalid expression"
// @Failure 500 {object} helpers.ErrorResponse "Error saving calculation"
// @Router /calculate [post]
func createCalculateHandler(calculator calculator) func(w http.ResponseWriter, r *http.Request) {
	var writeInternalErrorResponse = func(w http.ResponseWriter, err error) {
		log.Printf("Error when encoding get_history: %s", err)
		helpers.WriteErrorResponse(w, "something went wrong", http.StatusInternalServerError)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		calculation, err := calculator.PerformCalculationForUser(http_middleware.GetUserIDFromContext(r.Context()), req.Expression)
		if err != nil {
			if errors.Is(err, InvalidExpressionError) {
				helpers.WriteErrorResponse(w, "Invalid expression", http.StatusBadRequest)
				return
			}
			writeInternalErrorResponse(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(Response{Result: calculation.Result})
		if err != nil {
			writeInternalErrorResponse(w, err)
		}
	}
}
