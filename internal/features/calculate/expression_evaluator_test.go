package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGovaluateEvaluator_Evaluate(t *testing.T) {
	t.Run("returns result of valid expression", func(t *testing.T) {
		var (
			expressionString = "2 + 2"
			expectedResult   = 4.0
			evaluator        = &GovaluateEvaluator{}
		)

		result, err := evaluator.Evaluate(expressionString)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("returns error for invalid expression", func(t *testing.T) {
		var (
			expressionString = "invalidExpression"
			evaluator        = &GovaluateEvaluator{}
		)

		result, err := evaluator.Evaluate(expressionString)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, InvalidExpressionError, err)
	})
}
