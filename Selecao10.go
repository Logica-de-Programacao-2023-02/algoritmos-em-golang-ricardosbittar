package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Informe a idade:")
	fmt.Scanln(&idade)

	if idade <= 9 {
		fmt.Println("Mirim")
	}
	if idade >= 10 && idade <= 13 {
		fmt.Println("Infantil")
	}
	if idade >= 14 && idade <= 17 {
		fmt.Println("Juvenil")
	}
	if idade >= 18 {
		fmt.Println("Adulto")
	}
}
