package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	println(" Nova")
	println(" linha.")

	x := 3.14156
	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %f\n", x)
	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "Hey man!"

	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
