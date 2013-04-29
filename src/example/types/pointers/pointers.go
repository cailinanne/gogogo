package main
import "fmt"

type FancyInt int

type Vertex struct {
	X int
	Y int
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func zeroVertexVal(v Vertex){
	v.X = 0
	v.Y = 0
}

func zeroVertexPtr(v *Vertex){
	v.X = 0
	v.Y = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// New struct initialized with struct literal
	v1 := Vertex{1, 2}
	fmt.Println(v1)

	// Copies v1 to v2
	v2 := v1
	v2.X = 2
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("")

	// v3 is a pointer to v1
	v3 := &v1
	v3.X = 3
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println("")

	// This will modify only the copy of v1 which
	// is passed by value to the function
	zeroVertexVal(v1)
	fmt.Println(v1)
	fmt.Println("")

	// This will actually zero out v1 (and v3)
	zeroVertexPtr(&v1)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println("")

	// This returns a pointer to a Vertex
	v4 := new(Vertex)
	fmt.Println(v4)
	fmt.Println("")

	// Initializes a new FancyInt
	f1 := FancyInt(3)
	fmt.Println(f1)

	// Returns a pointer to a FancyInt
	// (which is weird and probably why you don't see this much)
	f2 := new(FancyInt)
	fmt.Println(f2)
	fmt.Println(*f2) // Dereference to display value

}
