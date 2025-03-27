package calculate

import (
	"OverEngineeredCalculator/database"
	"OverEngineeredCalculator/helpers"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

type calculationHistorySaver interface {
	Save(userID string, calculation database.Calculation) error
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
// @Router /api/v1/calculate [post]
func createCalculateHandler(calculationHistorySaver calculationHistorySaver) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		calculation, err := calculate(req.Expression)
		if err != nil {
			//TODO: consider adding different errorCode handling
			helpers.WriteErrorResponse(w, "Invalid expression", http.StatusBadRequest)
			return
		}

		userID := helpers.GetUserIDFromContext(r.Context())
		err = calculationHistorySaver.Save(userID, calculation)
		if err != nil {
			log.Printf("Error when saving calculation: %s", err)
			helpers.WriteErrorResponse(w, "something went wrong", http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(Response{Result: calculation.Result})
		if err != nil {
			log.Printf("Error when encoding get_history: %s", err)
			helpers.WriteErrorResponse(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}

// TODO: Extract to a calculate service and implement ... calculations ..
func calculate(expr string) (database.Calculation, error) {
	_, err := strconv.Atoi(expr)
	if err != nil {
		return database.Calculation{}, err
	}

	return database.Calculation{
		Expression: expr,
		Result:     expr,
	}, nil
}
