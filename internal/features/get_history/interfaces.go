package get_history

import "over-engineered-calculator/internal/database"

//go:generate go install github.com/matryer/moq@latest
//go:generate moq -out mocks_test.go . calculationHistoryRetriever

type calculationHistoryRetriever interface {
	GetHistory(userID string) ([]database.Calculation, error)
}
