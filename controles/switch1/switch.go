package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // executa o caso abaixo
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10.0))
	fmt.Println(notaParaConceito(9.4))
	fmt.Println(notaParaConceito(7.4))
	fmt.Println(notaParaConceito(5.3))
	fmt.Println(notaParaConceito(3.5))
	fmt.Println(notaParaConceito(1.5))
	fmt.Println(notaParaConceito(11.3))
}
