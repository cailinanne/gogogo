package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range tokens{
		// Relies on the fact that if the map does not contain the key, it returns
		// the appropriate "zero value" for the type
		m[value] = m[value] + 1
	}

	return m
}

func main() {
	fmt.Println(WordCount("i ate a donut then i ate another donut"))
}
