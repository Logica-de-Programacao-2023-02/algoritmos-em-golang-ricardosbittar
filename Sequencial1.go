package main

import "fmt"

func main() {

	var x int
	var y int
	var z int

	fmt.Println("Digite o primeiro numero: ")
	fmt.Scanln(&x)
	fmt.Println("Digite o segundo numero: ")
	fmt.Scanln(&y)
	fmt.Println("Digite o terceiro numero: ")
	fmt.Scanln(&z)
	soma := x + y + z
	fmt.Println("A soma dos tres numeros sao: ", soma)

}