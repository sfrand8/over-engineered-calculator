package calculate

import "over-engineered-calculator/internal/database"

//go:generate go install github.com/matryer/moq@latest
//go:generate moq -out mocks_test.go . calculator calculationHistorySaver expressionEvaluator

type expressionEvaluator interface {
	Evaluate(expressionString string) (interface{}, error)
}

type calculationHistorySaver interface {
	Save(userID string, calculation database.Calculation) error
}

type calculator interface {
	PerformCalculationForUser(userID string, expressionString string) (database.Calculation, error)
}
