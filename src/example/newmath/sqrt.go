package newmath

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot take square root of negative number: %f", float64(e))
}

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 0.0
	eps := 1.0
	for eps > 0.01 {
		zlast := z
		z -= (z*z - x) / (2 * x)
		eps = math.Abs(zlast - z)
	}
	return z, nil
}


