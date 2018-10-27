package main

import (
	"fmt"
)

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(8.5, 7.5, 8.0, 9.0))
	fmt.Printf("Média: %.2f\n", media(9.5, 9.0, 9.0, 9.0))
	fmt.Printf("Média: %.2f\n", media(10.0, 10.0, 10.0, 10.0))
}
