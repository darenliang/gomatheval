# A Simple Golang Math Expression Evaluator

[![GoDoc](https://godoc.org/github.com/darenliang/gomatheval?status.svg)](https://godoc.org/github.com/darenliang/go-math-eval)
[![Go report](http://goreportcard.com/badge/darenliang/gomatheval)](https://goreportcard.com/report/github.com/darenliang/go-math-eval)

A simple math expression evaluator written in Go.

To install:
```
go get github.com/darenliang/gomatheval
```

### Examples:
Evaluating expressions
```go
package main

import (
	"fmt"
	"github.com/darenliang/gomatheval"
	"math"
)

func main() {
	fmt.Println(gomatheval.EvalExpression(fmt.Sprintf("-(3+(sin(%v/2))^2)/4", math.Pi)))
	// outputs 1
}
```
Command-Line Application
```go
package main

import (
	"flag"
	"fmt"
	"github.com/darenliang/gomatheval"
)

func main() {
	optionPtr := flag.String("op", "e", "Select processing option.\ne: Evaluate expression\np: Process to Reverse Polish Notation\nt: Tokenize expression\n")
	expressionPtr := flag.String("exp", "", "Expression string to be processed")
	flag.Parse()
	switch *optionPtr {
	case "t":
		fmt.Println(gomatheval.SanitizeExpression(*expressionPtr))
	case "p":
		fmt.Println(gomatheval.ParseRPN(gomatheval.SanitizeExpression(*expressionPtr)))
	default:
		fmt.Println(gomatheval.EvalExpression(*expressionPtr))
	}
}
```

### What it supports:
* Basic operators such as `+ (addition)`, `- (subtraction)`, `* (multiplication)`, `/ (division)`, `^ (power)`.
* 64-bit floating point operations.
* Basic functions such as `sin (sine)`, `cos (cosine)`, `tan (tangent)`, `ln (natural logarithm)`, `lg (base 2 logarithm)`, `log (logarithm with variable base)`. More operators coming soon...
* Unary operators for positive and negative numbers. Includes proper handling of unary precedence operations.

### How it works:
Every expression is composed of three kinds of tokens:
* Operands
* Operators/Functions
* Parentheses

By tokenizing an expression, we can get the individual elements that make up the expression:
```
    -(3+(sin(3.141592653589793/2))^2)/4
=>  - ( 3 + ( sin ( 3.141592653589793 / 2 ) ) ^ 2 ) / 4
```

The tokenized expression is then sanitized and unary operators are converted to unique operators.
```
    - ( 3 + ( sin ( 3.141592653589793 / 2 ) ) ^ 2 ) / 4
=>  -u ( 3 + ( sin ( 3.141592653589793 / 2 ) ) ^ 2 ) / 4
```

By using the Shunting Yard Algorithm we can convert this infix notation into postfix notation (also known as Reverse Polish notation).
```
    -u ( 3 + ( sin ( 3.141592653589793 / 2 ) ) ^ 2 ) / 4
=>  3 3.141592653589793 2 / sin 2 ^ + -u 4 /
```

By evaluating the postfix notation tokens, we can obtain a final value.
```
    3 3.141592653589793 2 / sin 2 ^ + -u 4 /
=>  -1
```
