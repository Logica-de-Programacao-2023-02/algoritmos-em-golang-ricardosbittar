package main

import "fmt"

func main() {

	var x int

	fmt.Println("Digite um número: ")
	fmt.Scanln(&x)

	dobro := 2 * x
	triplo := 3 * x
	quadruplo := 4 * x

	fmt.Println("O dobro, o triplo e o quadruplo de", x, "sao:", dobro, triplo, quadruplo)
}
