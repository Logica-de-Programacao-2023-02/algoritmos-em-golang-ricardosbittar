package main

import "fmt"

func main() {
	var valor float32
	fmt.Println("valor do produto:")
	fmt.Scan(&valor)
	fmt.Println("o produto com 15% de desconto fica:", valor*0.85, "reais")
}
