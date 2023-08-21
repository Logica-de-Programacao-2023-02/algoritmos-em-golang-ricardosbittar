package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite um numero")
	fmt.Scanln(&x)

	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}
}
