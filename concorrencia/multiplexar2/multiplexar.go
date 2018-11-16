package main

import (
	"fmt"
	"time"
)

func fraseSecreta(frase string) <-chan string {
	canal := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			canal <- fmt.Sprintln(frase)
		}
	}()
	return canal
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				canal <- s
			case s := <-entrada2:
				canal <- s
			}
		}
	}()
	return canal
}

func main() {
	canal := juntar(fraseSecreta("Toda magia tem um preço"), fraseSecreta("queridinha hihihihi..."))

	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
}
