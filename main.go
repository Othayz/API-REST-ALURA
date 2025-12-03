package main

import "fmt"

type Pizza struct {
	id	int
	nome string
	preco float64
}

func main() {
	
	var nomePizzaria string = "Pizzaria GO"
	var toscana Pizza
	instagram, telefone := "@pizzaria_GO", 123456789
	
	fmt.Printf("Nome da pizzaria: %s (instagram: %s) - Telefone: %d", nomePizzaria, instagram, telefone)
}