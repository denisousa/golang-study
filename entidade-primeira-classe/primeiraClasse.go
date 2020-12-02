package main

import "fmt"

var soma = func(a, b, c int) int {
	return a + b + c
}

var testFunc = func(name string) string {
	return "yeah " + name
}

func main() {
	fmt.Println(soma(2, 3, 5))

	sub := func(a, b int) int {
		return a - b
	}

	testVar := 111
	fmt.Println(testVar)

	fmt.Println(testFunc("Denis"))

	fmt.Println(sub(2, 3))
}
