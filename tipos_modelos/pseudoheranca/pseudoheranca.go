package main

import (
	"fmt"
)

// Não é herança, e sim uma composição

type carro struct {
	nome            string
	velocidadeAtual int
}

type audi struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	a := audi{}
	a.nome = "Audi A4"
	a.velocidadeAtual = 0
	a.turboLigado = false

	fmt.Printf("O %s está com o turbo ligado? %v\n", a.nome, a.turboLigado)
	fmt.Println(a)
}
