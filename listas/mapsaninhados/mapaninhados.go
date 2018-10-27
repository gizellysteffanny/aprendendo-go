package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 12364.70,
			"Guga Pereira":  7346.57,
		},
		"J": {
			"José João": 23674.58,
		},
		"P": {
			"Pedro Junior": 2374.48,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
