package main

import (
	"fmt"
	"time"

	"github.com/gizellysteffanny/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	site1 := html.Titulo(url1)
	site2 := html.Titulo(url2)
	site3 := html.Titulo(url3)

	// estrutura de controle especifica para trabalhar com concorrência
	select {
	case t1 := <-site1:
		return t1
	case t2 := <-site2:
		return t2
	case t3 := <-site3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos Perderam!"
		// default:
		// 	return "Sem resposta ainda"
	}
}

func main() {
	vencedor := oMaisRapido(
		"https://www.github.com",
		"https://www.youtube.com",
		"https://www.google.com",
	)

	fmt.Println(vencedor)
}
