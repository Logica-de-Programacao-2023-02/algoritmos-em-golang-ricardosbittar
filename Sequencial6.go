package main

import (
	"fmt"
)

func main() {
	var salario float32
	fmt.Println("Digite seu salario: ")
	fmt.Scanln(&salario)
	novosalario := salario * 1.15
	fmt.Println("Seu salario com um aumento de 15 por cento ficou %2f: ", novosalario)
}
