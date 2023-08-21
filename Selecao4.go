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

	if imc < 18.5 {
		fmt.Println("Abaixo do peso normal")
	}
	if imc >= 18.5 && imc <= 24.9 {
		fmt.Println("Peso normal")
	}
	if imc >= 25.0 && imc <= 29.9 {
		fmt.Println("Excesso de peso")
	}
	if imc >= 30.0 && imc <= 34.9 {
		fmt.Println("Obesidade I")
	}
	if imc >= 35.0 && imc <= 39.9 {
		fmt.Println("Obesidade II")
	}
	if imc >= 40 {
		fmt.Println("Obesidade III")
	}
}
