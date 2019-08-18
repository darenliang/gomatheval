package gomatheval

import (
	"strconv"
	"strings"
	"unicode"
)

// SanitizeExpression returns interface of sanitized tokens ready to be parsed
func SanitizeExpression(expression string) []interface{} {
	var tokens = TokenizeExpression(expression)
	ParseFloats(&tokens)
	ProcessUnaryOperators(&tokens)
	return tokens
}

// TokenizeExpression returns interface of unprocessed tokens
func TokenizeExpression(expression string) []interface{} {
	var tokens []interface{}
	var currToken strings.Builder
	for _, char := range expression {
		if char != ' ' {
			if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '.' {
				currToken.WriteString(string(char))
			} else {
				var lastToken = currToken.String()
				if lastToken != "" {
					tokens = append(tokens, lastToken, string(char))
				} else {
					tokens = append(tokens, string(char))
				}
				currToken.Reset()
			}
		}
	}
	leftoverToken := currToken.String()
	if leftoverToken != "" && leftoverToken != " " {
		tokens = append(tokens, leftoverToken)
	}
	return tokens
}

// ParseFloats parses floats in tokens interface
func ParseFloats(tokens *[]interface{}) {
	for i, v := range *tokens {
		if strV, isStr := v.(string); isStr {
			if num, err := strconv.ParseFloat(strV, 64); err == nil {
				(*tokens)[i] = num
			}
		}
	}
}

// ProcessUnaryOperators processes unary operators in tokens
func ProcessUnaryOperators(tokens *[]interface{}) {
	for i := range *tokens {
		var currToken = (*tokens)[i]
		var isFloat bool
		if currToken == "-" || currToken == "+" {
			if i != 0 {
				_, isFloat = (*tokens)[i-1].(float64)
			}
			if i == 0 || !isFloat {
				switch currToken {
				case "-":
					(*tokens)[i] = "-u"
				case "+":
					(*tokens)[i] = "+u"
				}
			}
		}
	}
}
