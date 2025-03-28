package calculate

import (
	"github.com/go-chi/chi/v5"
	"over-engineered-calculator/internal/database"
)

func Setup(router chi.Router) {
	router.Post("/calculate", createCalculateHandler(database.GetCalculationHistoryRepository()))
}
