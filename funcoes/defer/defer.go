package main

import (
	"fmt"
)

func obterValorSalario(salario int) int {
	/*
		O defer vai executar antes do return
	*/
	defer fmt.Println("Fim!")
	if salario > 5000 {
		fmt.Println("Dev Sênior")
		return salario
	}
	fmt.Println("Dev Pleno ou Junior")
	return salario
}

func main() {
	fmt.Println(obterValorSalario(4000))
	fmt.Println(obterValorSalario(8000))
}
