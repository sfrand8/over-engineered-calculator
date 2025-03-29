package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGovaluateEvaluator_Evaluate(t *testing.T) {
	t.Run("returns result of valid expression", func(t *testing.T) {
		var (
			cases = []struct {
				expressionString string
				expectedResult   interface{}
			}{
				{
					expressionString: "1 + 1",
					expectedResult:   2.0,
				},
				{
					expressionString: "2 * 3",
					expectedResult:   6.0,
				},
				{
					expressionString: "4 / 2",
					expectedResult:   2.0,
				},
				{
					expressionString: "5 - 3",
					expectedResult:   2.0,
				},
				{
					expressionString: "2 + 3 * 4",
					expectedResult:   14.0,
				},
				{
					expressionString: "10 / 2 + 3",
					expectedResult:   8.0,
				},
				{
					expressionString: "2 * (3 + 4)",
					expectedResult:   14.0,
				},
				{
					expressionString: "2 + 3 * (4 - 1)",
					expectedResult:   11.0,
				},
			}
			evaluator = &GovaluateEvaluator{}
		)

		for _, s := range cases {
			t.Run(s.expressionString, func(t *testing.T) {
				result, err := evaluator.Evaluate(s.expressionString)

				assert.NoError(t, err)
				assert.Equal(t, s.expectedResult, result)
			})
		}
	})

	t.Run("returns error for invalid expression", func(t *testing.T) {
		var (
			expressionString = ""
			evaluator        = &GovaluateEvaluator{}
		)

		result, err := evaluator.Evaluate(expressionString)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, InvalidExpressionError, err)
	})
}
