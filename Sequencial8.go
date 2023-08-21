package main

import "fmt"

func main() {
	var diaria float32
	var dias float32

	fmt.Println("Qual o valor da sua diaria? ")
	fmt.Scanln(&diaria)
	fmt.Println("Quantos dias voce trabalhou? ")
	fmt.Scanln(&dias)
	salario := diaria * dias
	fmt.Println("Seu salario:", salario)

}
