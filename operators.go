package gomatheval

import "math"

// Map containing operator information
var operatorMap = map[string]struct {
	precedence         int
	rightAssociativity bool
	arity              int
	operation          func(float64, float64) float64
}{
	"+": {1, false, 2, func(a float64, b float64) float64 {
		return a + b
	}},
	"-": {1, false, 2, func(a float64, b float64) float64 {
		return a - b
	}},
	"*": {2, false, 2, func(a float64, b float64) float64 {
		return a * b
	}},
	"/": {2, false, 2, func(a float64, b float64) float64 {
		return a / b
	}},
	"^": {3, true, 2, func(a float64, b float64) float64 {
		return math.Pow(a, b)
	}},
	"ln": {4, true, 1, func(a float64, _ float64) float64 {
		return math.Log(a)
	}},
	"log": {4, true, 2, func(a float64, b float64) float64 {
		return math.Log(a) / math.Log(b)
	}},
	"sin": {4, true, 1, func(a float64, _ float64) float64 {
		return math.Sin(a)
	}},
	"cos": {4, true, 1, func(a float64, _ float64) float64 {
		return math.Cos(a)
	}},
	"tan": {4, true, 1, func(a float64, _ float64) float64 {
		return math.Tan(a)
	}},
	"-u": {5, true, 1, func(a float64, _ float64) float64 {
		return -a
	}},
	"+u": {5, true, 1, func(a float64, _ float64) float64 {
		return a
	}},
}
