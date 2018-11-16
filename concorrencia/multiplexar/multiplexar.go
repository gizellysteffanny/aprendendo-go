package main

import (
	"fmt"

	"github.com/gizellysteffanny/html"
)

/*
* Recebe um dado de um canal de origem e encaminha para um canal de destino
 */
func encaminhar(origem <-chan string, destino chan string) {
	for {
		// sempre que chegar dados do canal de origem, encaminha para o canal de destino
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go encaminhar(entrada1, canal)
	go encaminhar(entrada2, canal)
	return canal
}

func main() {
	canal := juntar(
		html.Titulo("https://www.github.com", "https://www.youtube.com"),
		html.Titulo("https://www.google.com", "https://www.slack.com"),
	)

	fmt.Println(<-canal, " | ", <-canal)
	fmt.Println(<-canal, " | ", <-canal)
}
