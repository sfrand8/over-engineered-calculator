package get_history

import (
	"github.com/go-chi/chi/v5"
	"over-engineered-calculator/internal/database"
)

func Setup(mux chi.Router) {
	mux.Get("/get_history", createHistoryHandler(database.GetCalculationHistoryRepository()))
}
