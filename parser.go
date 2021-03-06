package gomatheval

// ParseRPN parses tokens to RPN
func ParseRPN(tokens []interface{}) []interface{} {
	var stack = expressionStack{}
	for _, token := range tokens {
		switch token {
		case "(":
			stack.operators = append(stack.operators, token.(string))
		case ")":
			for stack.operators[len(stack.operators)-1] != "(" {
				stack.popAndAppend()
			}
			stack.operators = stack.operators[:len(stack.operators)-1]
		default:
			if strToken, isStr := token.(string); isStr {
				tokenOp := operatorMap[strToken]
				for len(stack.operators) != 0 {
					if topOp, isOp := operatorMap[stack.operators[len(stack.operators)-1]]; !isOp ||
						tokenOp.precedence > topOp.precedence ||
						tokenOp.precedence == topOp.precedence && tokenOp.rightAssociativity ||
						tokenOp.rightAssociativity && topOp.arity == 1 {
						break
					}
					stack.popAndAppend()
				}
				stack.operators = append(stack.operators, strToken)
			} else {
				stack.output = append(stack.output, token)
			}
		}
	}
	for len(stack.operators) != 0 {
		stack.popAndAppend()
	}
	return stack.output
}
