package main

import "fmt"

func main() {
	var x int

	fmt.Println("Digite um numero")
	fmt.Scanln(&x)

	for i := 0; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}

}
