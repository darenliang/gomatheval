package go_math_eval

func (stack *expressionStack) popAndAppend() {
	stack.output, stack.operators = append(stack.output, stack.operators[len(stack.operators)-1]), stack.operators[:len(stack.operators)-1]
}

type expressionStack struct {
	operators []string
	output    []interface{}
}
