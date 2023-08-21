package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite um número:")
	fmt.Scanln(&x)
	fmt.Println("O antecessor e sucessor de", x, "sao", x-1, "e", x+1)

}
