package main

import "fmt"

func checkNote(n float32) string {
	switch {
	case n >= 9:
		return "A"
	case n < 9 && n >= 7:
		return "B"
	case n < 7 && n > 4:
		return "C"
	case n <= 4:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(checkNote(2.0))
	fmt.Println(checkNote(10.0))
	fmt.Println(checkNote(7.0))
}
