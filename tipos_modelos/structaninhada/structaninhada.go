package main

import (
	"fmt"
)

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item // slice de item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 5, preco: 2.50},
			item{produtoID: 6, quantidade: 8, preco: 1.75},
			item{4, 10, 0.99}, // evite usar a struct em ordem, fica confuso
		},
	}

	fmt.Printf("O valor total do pedido é de R$ %.2f", pedido.valorTotal())
}
