package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador conta
	fmt.Println(numeros)

	for index, element := range numeros {
		fmt.Printf("%d) %d\n", index+1, element)
	}

	for _, elemento := range numeros { // o underline ignora o index
		fmt.Println(elemento)
	}
}
