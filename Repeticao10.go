package main

import "fmt"

func main() {
	var x int
	var maior int
	for {
		fmt.Println("Digite um numero: ")
		fmt.Scanln(&x)

		if x == 0 {
			break
		}
		if x > maior {
			maior = x
		}
	}
	if maior != 0 {
		fmt.Println("O maior numero é:", maior)
	} else {
		fmt.Println("Voce precisa inserir um numero")
	}

}
