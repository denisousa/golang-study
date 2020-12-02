package main

import "fmt"

func doubleValue(arg *int) { // Pointe of int
	fmt.Println(*arg)
	*arg++
	fmt.Println(*arg)
}

func main() {
	fmt.Println("Ponteiros")

	n := 9
	doubleValue(&n)

}
