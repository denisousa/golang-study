package main

import "fmt"

func f1() {
	fmt.Println("I am F1")
}

func f2() string {
	return "I am F2"
}

func f3(param string) {
	fmt.Println("I am " + param)
}

func f4(param string) string {
	// Sprintf -> return string
	return fmt.Sprintf("I am %s", param)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()              // Return empity
	fmt.Println(f2()) // Return string
	f3("F3")          // Return empity and receive param
	fmt.Println(f4("F4"))
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
