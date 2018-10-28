package main

import (
	"fmt"
)

func incremeta(n int) {
	n++
}

// Um ponteiro é representado por um *
func decrementa(n *int) {
	// O * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n--
}

func main() {
	n := 4

	incremeta(n) // por valor
	fmt.Println(n)

	// O & é usado para obter o endereço da variável
	decrementa(&n) // por referência
	fmt.Println(n)

	/*
		Sempre tentar evitar a passagem de referência para outros métodos, pois pode gerar efeitos colaterais, ou seja,
		a função @decrementa deixou de ser uma função pura, ela trabalha com coisas que são externas...

		Priorizar sempre a passagem por valor...
	*/

}
