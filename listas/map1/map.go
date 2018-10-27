package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Rick"
	aprovados[747537857] = "Nick"
	aprovados[906896895] = "Dick"
	aprovados[367427858] = "Dawn"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[747537857])
	delete(aprovados, 123456789)
	fmt.Println(aprovados[123456789])
	fmt.Println(aprovados)

}
