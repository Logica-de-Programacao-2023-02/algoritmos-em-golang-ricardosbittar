package main

import "fmt"

func main() {
	var x int
	var soma int
	var contador int

	for {
		fmt.Println("Digite um número, e digite 0 para interromper")
		fmt.Scanln(&x)

		if x == 0 {
			break
		}
		soma += x
		contador++
	}
	if contador > 0 {
		media := float64(soma) / float64(x)
		fmt.Println("média aritmética: %.2f", media)
	} else {
		fmt.Println("Nenhum número inserido")
	}

}
