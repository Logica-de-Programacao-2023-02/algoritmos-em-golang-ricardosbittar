package main

import "fmt"

func main() {
	var peso float32

	fmt.Println("Digite seu peso em quilos:")
	fmt.Scan(&peso)
	fmt.Println("Seu peso em libras é", peso/0.45)
}
