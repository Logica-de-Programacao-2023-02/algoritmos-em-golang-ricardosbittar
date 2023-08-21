package main

import "fmt"

func main() {
	var salario float32
	fmt.Println("salario: ")
	fmt.Scanln(&salario)

	if salario < 1000.00 {
		fmt.Println("Salario com 10% de aumento:", salario*1.10)
	}
	if salario >= 1000.00 {
		fmt.Println("Salario com 5% de aumento:", salario*1.05)
	}
}
