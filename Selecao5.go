package main

import "fmt"

func main() {
	var x int
	fmt.Println("numero:")
	fmt.Scanln(&x)

	if x%3 == 0 && x%^5 == 0 {
		fmt.Println(x, "é multiplo de 3 e 5 ao mesmo tempo")
	} else {
		fmt.Println(x, "nao é multiplo de 3 e 5 ao mesmo tempo")
	}
}
