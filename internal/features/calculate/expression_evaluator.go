package calculate

import "github.com/Knetic/govaluate"

// GovaluateEvaluator is a struct created to hide the external library, so I can mock it in tests of my calculator
type GovaluateEvaluator struct{}

func (g *GovaluateEvaluator) Evaluate(expressionString string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(expressionString)
	if err != nil {
		return nil, InvalidExpressionError
	}

	return expression.Evaluate(nil)
}
