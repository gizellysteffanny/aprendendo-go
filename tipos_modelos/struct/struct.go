package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Mouse sem fio",
		preco:    58.30,
		desconto: 0.05,
	}
	fmt.Println(produto1)
	fmt.Printf("Preço com desconto: %.2f\n", produto1.precoComDesconto())

	// forma curta de passa os valores da struct produto
	// para funcionar os valores tem que ser passado na ordem corrente
	produto2 := produto{"Teclado Gamer", 149.99, 0.47}
	fmt.Println(produto2)
	fmt.Printf("Preço com desconto: %.2f\n", produto2.precoComDesconto())
}
