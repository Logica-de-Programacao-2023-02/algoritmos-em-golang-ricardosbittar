package main

import "fmt"

func main() {

	var peso float32
	var altura float32

	fmt.Println("Qual o seu peso? ")
	fmt.Scanln(&peso)
	fmt.Println("Qual a sua altura? ")
	fmt.Scanln(&altura)
	imc := peso / altura * altura
	fmt.Println("O seu IMC é:", imc)

}
