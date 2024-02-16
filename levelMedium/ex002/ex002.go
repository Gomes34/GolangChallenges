package main

import "fmt"

type documentos struct {
	nome   string
	idade  int
	altura int
	casada bool
}

func main() {

	pessoa := documentos{
		nome:   "Euler",
		idade:  19,
		altura: 177,
		casada: false,
	}

	fmt.Println(pessoa)
	fmt.Println(pessoa.altura)

}
