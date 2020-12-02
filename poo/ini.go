package main

import "fmt"

type product struct {
	name     string
	price    float32
	discount float32
}

func (p product) applyDiscount() float32 {
	return p.price * (1 - p.discount)
}

func main() {
	var pencil product
	pencil = product{
		name:     "pencil",
		price:    4.0,
		discount: 0.01,
	}

	fmt.Println(pencil, pencil.applyDiscount())

	notebook := product{"Notebook", 1989.90, 0.10}
	fmt.Println(notebook, notebook.applyDiscount())
}
