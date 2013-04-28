/**
 * Created with IntelliJ IDEA.
 * User: cailinnelson
 * Date: 4/28/13
 * Time: 8:34 AM
 * To change this template use File | Settings | File Templates.
 */
package newmath

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	// This is a terrible implementation.
	// Real code should import "math" and use math.Sqrt.
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}


