package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, canal chan int) {
	time.Sleep(time.Second)
	canal <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	canal <- 3 * base

	time.Sleep(3 * time.Second)
	canal <- 4 * base
}

func main() {
	canal := make(chan int)
	go doisTresQuatroVezes(2, canal)

	a, b := <-canal, <-canal // recebendo dados do canal

	fmt.Println(a, b)

	fmt.Println(<-canal)
}
