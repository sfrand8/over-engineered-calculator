package get_history

import (
	"OverEngineeredCalculator/database"
	"github.com/go-chi/chi/v5"
)

func Setup(mux chi.Router) {
	mux.Get("/get_history", createHistoryHandler(database.GetCalculationHistoryRepository()))
}
