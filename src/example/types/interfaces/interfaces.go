package main
import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func saySomething(a Animal){
	fmt.Println(a.Speak())
}

func main() {

	// Inside the array you must use a pointer to a Cat because the
	// method speak is implemented on the pointer.
	animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	animals2 := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}

	for _, animal := range animals2 {
		fmt.Println(animal.Speak())
	}

	// However.... if you have just one Cat outside of an array you
	// can do it either way?
	cat1 := Cat{}
	fmt.Println(cat1)
	fmt.Println(cat1.Speak())

	cat2 := new(Cat)
	fmt.Println(cat2)
	fmt.Println(cat2.Speak())

	// If you pass just one Cat to a method then it only works if
	// you pass the pointer
	saySomething(&cat1)
	saySomething(cat2)


}
