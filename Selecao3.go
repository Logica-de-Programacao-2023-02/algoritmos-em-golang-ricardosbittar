package main

import "fmt"

func main() {
	var x int
	fmt.Println("Numero: ")
	fmt.Scanln(&x)
	if x%2 == 0 {
		fmt.Println(x, "é par")
	} else {
		fmt.Println(x, "é ímpar")

	}

}
