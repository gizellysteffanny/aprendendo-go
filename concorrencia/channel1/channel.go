package main

import (
	"fmt"
)

func main() {
	canal := make(chan int, 1)
	
	canal <- 1 // enviando dados para o canal (escrita)

	<-canal // recebendo dados do canal (leitura)

	canal <- 2
	fmt.Println(<-canal)
}