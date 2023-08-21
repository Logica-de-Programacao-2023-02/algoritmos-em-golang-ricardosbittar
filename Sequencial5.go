package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Qual a idade? ")
	fmt.Scanln(&idade)
	dias := idade * 365
	fmt.Println("A SUA IDADE EM DIAS É:", dias)

}
