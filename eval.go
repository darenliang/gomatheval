package gomatheval

// EvalRPN evaluates tokens in RPN to float64
func EvalRPN(tokens []interface{}) float64 {
	var numStack = make([]float64, 0)
	for _, token := range tokens {
		switch value := token.(type) {
		case string:
			var operandSlice []float64
			operandSlice = append(operandSlice, numStack[len(numStack)-1])
			switch operatorMap[value].arity {
			case 2:
				operandSlice = append(operandSlice, numStack[len(numStack)-2])
				numStack[len(numStack)-2] = operatorMap[value].operation(operandSlice)
				numStack = numStack[:len(numStack)-1]
			default:
				numStack[len(numStack)-1] = operatorMap[value].operation(operandSlice)
			}
		default:
			if num, isNum := value.(float64); isNum {
				numStack = append(numStack, num)
			}
		}
	}
	return numStack[0]
}

// EvalExpression evaluates expression in string to float64
func EvalExpression(expression string) float64 {
	return EvalRPN(ParseRPN(SanitizeExpression(expression)))
}
