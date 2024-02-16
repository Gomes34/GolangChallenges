package main

import "fmt"

func main() {

	dev := struct {
		nome  string
		idade int
	}{
		nome:  "Euler",
		idade: 19,
	}

	fmt.Println(dev)

}
