package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range tokens{
		m[value] = m[value] + 1
	}

	return m
}

func main() {
	fmt.Println(WordCount("i ate a donut then i ate another donut"))
}
