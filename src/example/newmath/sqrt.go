/**
 * Created with IntelliJ IDEA.
 * User: cailinnelson
 * Date: 4/28/13
 * Time: 8:34 AM
 * To change this template use File | Settings | File Templates.
 */
package newmath

import (
	"fmt"
	"math"
)

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	// This is a terrible implementation.
	// Real code should import "math" and use math.Sqrt.
	z := 0.0
	eps := 1.0
	for eps > 0.01 {
		zlast := z
		z -= (z*z - x) / (2 * x)
		eps = math.Abs(zlast - z)
		fmt.Println(eps)

	}
	return z
}


