package main

import (
	"fmt"
)

func main() {
	funcionariosESalarios := map[string]float64{
		"Rick": 1500.47,
		"Nick": 8000.50,
		"Dick": 2600.35,
		"Dawn": 1200.12,
	}

	funcionariosESalarios["MilHouse"] = 1900.17
	delete(funcionariosESalarios, "Inexistente")

	for nome, salario := range funcionariosESalarios {
		fmt.Println(nome, salario)
	}

}
