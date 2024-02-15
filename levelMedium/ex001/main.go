package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	empregada bool
	lugar     string
}

func main() {
	pessoa1 := profissional{
		pessoa: pessoa{
			nome:  "Euler",
			idade: 19,
		},
		empregada: false,
		lugar:     "Camaqua",
	}

	if pessoa1.empregada {
		fmt.Printf("Olá %v, você está empregado em %v", pessoa1.nome, pessoa1.lugar)
	} else {
		fmt.Printf("Olá %v, você não está empregado em %v", pessoa1.nome, pessoa1.lugar)
	}
}
