package gomatheval

// EvalRPN evaluates tokens in RPN to float64
func EvalRPN(tokens []interface{}) float64 {
	var numStack = make([]float64, 0)
	for _, token := range tokens {
		switch value := token.(type) {
		case string:
			var lastNum = &numStack[len(numStack)-1]
			switch operatorMap[value].arity {
			case 2:
				var secondLastNum = &numStack[len(numStack)-2]
				*secondLastNum = operatorMap[value].operation(*secondLastNum, *lastNum)
				numStack = numStack[:len(numStack)-1]
			default:
				*lastNum = operatorMap[value].operation(*lastNum, 0)
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
