package go_math_eval

func (stack *ExpressionStack) popAndAppend() {
	stack.output, stack.operators = append(stack.output, stack.operators[len(stack.operators)-1]), stack.operators[:len(stack.operators)-1]
}

type ExpressionStack struct {
	operators []string
	output    []interface{}
}
