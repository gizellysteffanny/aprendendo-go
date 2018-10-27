package main

import (
	"fmt"
)

func main() {
	array := [3]int{1, 2, 3}
	var slice1 []int

	// array = append(array, 4,5,6) => Não é permitido no array

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
