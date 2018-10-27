package main

import (
	"fmt"
)

func imprimirContos(contos ...string) {
	fmt.Println("Era uma vez...")
	for index, conto := range contos {
		fmt.Printf("%d) %s\n", index+1, conto)
	}
}

func main() {
	contos := []string{"João e Maria", "Branca de Neve", "Bela e a Fera", "Cinderela", "Chapeuzinho vermelho"}

	imprimirContos(contos...)
}
