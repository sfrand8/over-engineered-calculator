package calculate

import (
	"fmt"
	"over-engineered-calculator/internal/database"
)

// CalculatorService is a struct that contains an expressions evaluator and returns the result of the expression after saving it to the historySaver of the user
// I didn't think the important aspect of this code challenge was to see me create a calculator
// but rather how I would structure the code to be maintainable and extensible
// So I just decided to use a library that does the heavy lifting for me and spend time on the structure of the code.
// Had I needed to implement an actual calculator I would have looked at the string and treated it as a compiler treats any other language
// I would have created a lexer to tokenize the string and then a parser to parse the tokens into an abstract syntax tree
// Then I would have traversed the tree and evaluated it in a recursive manner ending up with the final result in the end at the root node
type CalculatorService struct {
	evaluator    expressionEvaluator
	historySaver calculationHistorySaver
}

func NewCalculator(evaluator expressionEvaluator, history calculationHistorySaver) *CalculatorService {
	return &CalculatorService{evaluator: evaluator, historySaver: history}
}

// PerformCalculationForUser processes an expression string and returns the calculation result
func (c *CalculatorService) PerformCalculationForUser(userID string, expressionString string) (database.Calculation, error) {
	result, err := c.evaluator.Evaluate(expressionString)
	if err != nil {
		return database.Calculation{}, err
	}

	calculation := database.Calculation{
		Expression: expressionString,
		Result:     fmt.Sprintf("%v", result),
	}

	err = c.historySaver.Save(userID, calculation)
	if err != nil {
		return database.Calculation{}, err
	}

	return calculation, nil
}
