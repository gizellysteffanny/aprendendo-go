package main

import (
	"fmt"
)

/*
Provando que o slice é uma estrutura que tem: tamanho e um ponteiro para o elemento de um array
*/
func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	// Muda o valor nos dois slices, pois eles aponta para o mesmo array
	s1[0] = 3
	fmt.Println(s1, s2)
}
