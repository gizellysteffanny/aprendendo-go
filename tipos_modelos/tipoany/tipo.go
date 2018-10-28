package main

import (
	"fmt"
)

type curso struct {
	nome string
}

func main() {
	var angular interface{}
	fmt.Println(angular)

	angular = "Angular 6"
	fmt.Println(angular)

	angular = 29.99
	fmt.Println(angular)

	type any interface{}

	var rick any = "Gato loiro"
	fmt.Println(rick)

	rick = true
	fmt.Println(rick)
}
