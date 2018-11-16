package main

import (
	"fmt"
	"time"
)

func rotina(canal chan int) {
	canal <- 1
	canal <- 2
	canal <- 3
	fmt.Println("Executou !")
	canal <- 4
	canal <- 5
	canal <- 6
}

func main() {
	canal := make(chan int, 3) // 3 buffers
	go rotina(canal)

	time.Sleep(time.Second)
	println(<-canal)
	println(<-canal)
}
