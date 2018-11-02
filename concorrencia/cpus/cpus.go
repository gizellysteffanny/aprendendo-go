package main

import (
	"fmt"
	"runtime"
)

func main() {
	// mostra quantos cpus estão rodando na máquina
	fmt.Println(runtime.NumCPU())
}
