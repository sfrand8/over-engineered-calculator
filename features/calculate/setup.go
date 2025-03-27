package calculate

import (
	"OverEngineeredCalculator/database"
	"github.com/go-chi/chi/v5"
)

func Setup(router chi.Router) {
	router.Post("/calculate", createCalculateHandler(database.GetCalculationHistoryRepository()))
}
