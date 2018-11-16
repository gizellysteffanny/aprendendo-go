package main

import (
	"fmt"
	"time"
)

func rotina(canal chan int) {
	time.Sleep(time.Second)
	canal <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c)   // deadlock
	fmt.Println("Fim") // nunca vai ser impresso pq em cima gerou um deadlock
}
