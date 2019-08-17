package gomatheval

import (
	"strconv"
	"strings"
	"unicode"
)

// Sanitizes expression to tokens ready to be parsed
func SanitizeExpression(expression string) []interface{} {
	var tokens = TokenizeExpression(expression)
	ParseFloats(&tokens)
	ProcessUnaryOperators(&tokens)
	return tokens
}

// Tokenize expressions to string tokens
func TokenizeExpression(expression string) []interface{} {
	var tokens []interface{}
	var currToken strings.Builder
	for _, chr := range expression {
		if chr != ' ' {
			if unicode.IsDigit(chr) || unicode.IsLetter(chr) || chr == '.' {
				currToken.WriteString(string(chr))
			} else {
				var lastToken = currToken.String()
				if lastToken != "" {
					tokens = append(tokens, lastToken, string(chr))
				} else {
					tokens = append(tokens, string(chr))
				}
				currToken.Reset()
			}
		}
	}
	leftOverToken := currToken.String()
	if leftOverToken != "" {
		tokens = append(tokens, leftOverToken)
	}
	return tokens
}

// Parse floats in tokens
func ParseFloats(tokens *[]interface{}) {
	for i, v := range *tokens {
		if strV, isStr := v.(string); isStr {
			if num, err := strconv.ParseFloat(strV, 64); err == nil {
				(*tokens)[i] = num
			}
		}
	}
}

// Process unary operators in tokens
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
