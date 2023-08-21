package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	peso1 := 2
	peso2 := 3
	peso3 := 5
	fmt.Println("Qual o primeiro número? ")
	fmt.Scanln(&x)
	fmt.Println("Qual o segundo número? ")
	fmt.Scanln(&y)
	fmt.Println("Qual o terceiro número? ")
	fmt.Scanln(&z)
	media := x*peso1 + y*peso2 + z*peso3/peso1 + peso2 + peso3
	fmt.Println("A média Ponderada é", media)
}
