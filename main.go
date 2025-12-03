package main

import "fmt"

type Pizza struct {
	id	int
	nome string
	preco float64
}

func main() {
	
	var pizzas = []Pizza{
		{id: 1, nome: "Toscana", preco: 25.00},
		{id: 2, nome: "Marquerita", preco: 30.00},
		{id: 3, nome: "Atum com queijo", preco: 35.00},
	}
	fmt.Println(pizzas)	
}