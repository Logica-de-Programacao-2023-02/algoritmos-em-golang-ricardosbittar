package main

func main() {
lista := [6, 8, 5, 7, 1, 9]
menor := lista[0]

for _, num:= range lista{
	if num < menor {
		menor = num
}

	}
 fmt.Println("O menor numero é", menor)
}