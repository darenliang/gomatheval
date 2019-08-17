package gomatheval

import "math"

// operatorMap contains a map of operator information
var operatorMap = map[string]struct {
	precedence         int
	rightAssociativity bool
	arity              int
	operation          func([]float64) float64
}{
	"+": {1, false, 2, func(params []float64) float64 {
		return params[1] + params[0]
	}},
	"-": {1, false, 2, func(params []float64) float64 {
		return params[1] - params[0]
	}},
	"*": {2, false, 2, func(params []float64) float64 {
		return params[1] * params[0]
	}},
	"/": {2, false, 2, func(params []float64) float64 {
		return params[1] / params[0]
	}},
	"^": {3, true, 2, func(params []float64) float64 {
		return math.Pow(params[1], params[0])
	}},
	"ln": {4, true, 1, func(params []float64) float64 {
		return math.Log(params[0])
	}},
	"log": {4, true, 2, func(params []float64) float64 {
		return math.Log(params[1]) / math.Log(params[0])
	}},
	"sin": {4, true, 1, func(params []float64) float64 {
		return math.Sin(params[0])
	}},
	"cos": {4, true, 1, func(params []float64) float64 {
		return math.Cos(params[0])
	}},
	"tan": {4, true, 1, func(params []float64) float64 {
		return math.Tan(params[0])
	}},
	"-u": {5, true, 1, func(params []float64) float64 {
		return -params[0]
	}},
	"+u": {5, true, 1, func(params []float64) float64 {
		return +params[0]
	}},
	"abs": {5, true, 1, func(params []float64) float64 {
		return math.Abs(params[0])
	}},
}
