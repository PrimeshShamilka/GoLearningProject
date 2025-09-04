package testing

import (
	"errors"
)

// common errors
var (
	ErrNegSqrt = errors.New("cannot Sqrt negative number")
	ErrNoSolution = errors.New("no solution found")
)

// returns the absolute value of a float64
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// computes the square root of a number using Newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegSqrt
	}
	z := x / 2.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}