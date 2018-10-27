package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case bool:
		return "Booleano"
	case func():
		return "Função"
	default:
		return "Tipo não detectado"
	}
}

func main() {
	fmt.Println(tipo(3))
	fmt.Println(tipo(9.5))
	fmt.Println(tipo("Hey man!!!"))
	fmt.Println(tipo(false))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
