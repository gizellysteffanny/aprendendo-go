package main

import (
	"fmt"
)

/*
Um closure (fechamento) é uma função que se "lembra" do ambiente — ou escopo léxico — em que ela foi criada.
*/

func closure() func() {
	x := 10
	// essa função anonima é uma função aninhada(um closure)
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

/*
A função não perde seu valor,
mesmo tendo uma variável com o mesmo nome na função main
*/
func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
