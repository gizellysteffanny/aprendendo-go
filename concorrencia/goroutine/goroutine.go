package main

import (
	"fmt"
	"time"
)

func fale(pessoa, text string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, text, i+1)
	}
}

func main() {
	fale("Rick", "I'm Rick", 5)
	fale("Nick", "I'm Nick", 8)

	go fale("Dick", "I'm Dick", 15)
	go fale("Fred", "I'm Fred", 35)

	go fale("Dawn", "I'm Dawn", 13)
	fale("Rumperstilskin", "I'm Nick", 3)
}
