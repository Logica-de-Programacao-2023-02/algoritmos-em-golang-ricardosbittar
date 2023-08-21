package main

import (
	"fmt"
	"sort"
)

func main() {
	var lista []int

	for i := 1; i <= 3; i++ {
		var num int
		fmt.Printf("Digite o número %d: ", i)
		fmt.Scan(&num)
		lista = append(lista, num)
	}

	sort.Ints(lista)

	fmt.Println("Números em ordem crescente:", lista)

}
