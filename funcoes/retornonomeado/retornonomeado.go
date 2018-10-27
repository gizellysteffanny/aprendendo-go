package main

import (
	"fmt"
)

func trocaDeValores(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo
}

func main() {
	x, y := trocaDeValores(2, 3)
	fmt.Println(x, y)

	primeiro, segundo := trocaDeValores(8, 5)
	fmt.Println(primeiro, segundo)
}
