package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// adicionar mais métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbooooo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	var bmw esportivoLuxuoso = bmw7{}
	bmw.ligarTurbo()
	bmw.fazerBaliza()
}
