# Just Another Golang Math Expression Evaluator
This is just another golang math expression evaluator written in pure Go.

### What it supports:
* Basic operators such as `+ (addition)`, `- (subtraction)`, `* (multiplication)`, `/ (division)`, `^ (power)`.
* 64-bit floating point operations.
* Basic functions such as `sin (sine)`, `cos (cosine)`, `tan (tangent)`, `ln (natural logarithm )`, `log (logarithm with variable base)`. More operators comming soon...
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