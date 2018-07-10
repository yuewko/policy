package policy

import (
	"strings"

	"context"
	"github.com/Knetic/govaluate"
)

type policyRule struct {
	keyword string
	expression  []string
}

type policyList struct {
	ruleList []policyRule
}

// validateExpression checks whether the format of the expression used is valid or not.
func validateExpression(args []string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(strings.Join(args, " "))
	if err != nil {
		return nil, err
	}

	parameters := make(map[string]interface{}, 8)
	// arg[0] will be either query or reply
	// arg[1] will be the variable which will be taken from the metadata plugin
	parameters[args[1]] = ""

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return nil, err
	}

	switch result {
	case true:
	case false:

	}

	return result, nil
}

func evaluateRule(p *policyList, ctx context.Context) error {

	switch rule {
	case typeAllow:
	case typeRefuse:

	default:

	}

	return nil
}
