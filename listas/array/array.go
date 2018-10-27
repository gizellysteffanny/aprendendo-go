package main

import (
	"fmt"
)

func main() {
	// homogênea (mesmo tipo) e estática(fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 8.5, 7.2, 9.5
	// notas[3] = 6.5 <= o array se torna inválido pois passou do tamanho permitido
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média: %.1f\n", media)
}
