package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite um numero entre 1 e 7:")
	fmt.Scanln(&x)

	if x == 1 {
		fmt.Println("Domingo")
	}
	if x == 2 {
		fmt.Println("Segunda")
	}
	if x == 3 {
		fmt.Println("terca")
	}
	if x == 4 {
		fmt.Println("Quarta")
	}
	if x == 5 {
		fmt.Println("Quinta")
	}
	if x == 6 {
		fmt.Println("Sexta")
	}
	if x == 7 {
		fmt.Println("Sabado")
	}
}
