package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Primeiro numero: ")
	fmt.Scanln(&x)
	fmt.Println("Segundo numero:")
	fmt.Scanln(&y)

	if x >= 0 && y >= 0 {
		fmt.Println(x * y)
	}
	if x < 0 || y < 0 {
		fmt.Println(x + y)

	}
}
