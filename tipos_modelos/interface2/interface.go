package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type audi struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (a *audi) ligarTurbo() {
	a.turboLigado = true
}

func main() {
	audiA4 := audi{"Audi A4", false, 0}
	audiA4.ligarTurbo()

	var audiA6 esportivo = &audi{"Audi A6", false, 0}
	audiA6.ligarTurbo()

	fmt.Println(audiA4, audiA6)
}
