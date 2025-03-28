package calculate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"over-engineered-calculator/internal/database"
	"testing"
)

func TestCalculatorService_PerformCalculationForUser(t *testing.T) {
	t.Run("returns result of calculation when everything succeeds", func(t *testing.T) {
		var (
			expectedResult      = 2.34
			expressionString    = "someExpression"
			userID              = "someUserID"
			expectedCalculation = database.Calculation{
				Expression: expressionString,
				Result:     fmt.Sprintf("%v", expectedResult),
			}
			mockEvaluator = &expressionEvaluatorMock{
				EvaluateFunc: func(expressionString string) (interface{}, error) {
					return expectedResult, nil
				},
			}
			mockHistorySaver = &calculationHistorySaverMock{
				SaveFunc: func(userID string, calculation database.Calculation) error {
					return nil
				},
			}
			sut = NewCalculator(mockEvaluator, mockHistorySaver)
		)

		result, err := sut.PerformCalculationForUser(userID, expressionString)

		assert.NoError(t, err)
		assert.Equal(t, expectedCalculation, result)
	})

	t.Run("returns error when evaluator fails", func(t *testing.T) {
		var (
			expressionString = "someExpression"
			userID           = "someUserID"
			mockEvaluator    = &expressionEvaluatorMock{
				EvaluateFunc: func(expressionString string) (interface{}, error) {
					return nil, assert.AnError
				},
			}
			mockHistorySaver = &calculationHistorySaverMock{}
			sut              = NewCalculator(mockEvaluator, mockHistorySaver)
		)

		result, err := sut.PerformCalculationForUser(userID, expressionString)

		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, database.Calculation{}, result)
		assert.Len(t, mockEvaluator.EvaluateCalls(), 1)
		assert.Equal(t, mockEvaluator.EvaluateCalls()[0].ExpressionString, expressionString)
	})

	t.Run("returns error when history saver fails", func(t *testing.T) {
		var (
			expectedResult      = 2.34
			expressionString    = "someExpression"
			userID              = "someUserID"
			expectedCalculation = database.Calculation{
				Expression: expressionString,
				Result:     fmt.Sprintf("%v", expectedResult),
			}
			mockEvaluator = &expressionEvaluatorMock{
				EvaluateFunc: func(expressionString string) (interface{}, error) {
					return expectedResult, nil
				},
			}
			mockHistorySaver = &calculationHistorySaverMock{
				SaveFunc: func(userID string, calculation database.Calculation) error {
					return assert.AnError
				},
			}
			sut = NewCalculator(mockEvaluator, mockHistorySaver)
		)

		result, err := sut.PerformCalculationForUser(userID, expressionString)

		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, database.Calculation{}, result)
		assert.Len(t, mockHistorySaver.SaveCalls(), 1)
		assert.Equal(t, mockHistorySaver.SaveCalls()[0].UserID, userID)
		assert.Equal(t, mockHistorySaver.SaveCalls()[0].Calculation, expectedCalculation)
	})
}
