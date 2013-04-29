package main
import "fmt"

type rect struct {
	width, height int
}

// Defined on the pointer so that it can modify the
// internals of the struct
func (r *rect) scale(scale int) {
	r.width *= scale
	r.height *= scale
}

func (r *rect) area() int {
	return r.width * r.height
}

// Defined on the value - cannot modify the internals
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println("")

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("")

	r.scale(2)
	fmt.Println("area: ", r.area())
	fmt.Println("area: ", rp.area())

	rp.scale(2)
	fmt.Println("area: ", r.area())
	fmt.Println("area: ", rp.area())

}
