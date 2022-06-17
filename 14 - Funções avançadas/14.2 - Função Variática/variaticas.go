package main

import "fmt"

//Passa quantos numeros quiser!
func soma(numeros ...int) int {
	total := 0
	for _, numeros := range numeros {
		total += numeros
	}
	return total
}

func escrever(texto string, numeros...int)  {
	for _, numeros := range numeros {
		fmt.Println(texto, numeros)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 103232)
	fmt.Println(totalDaSoma)	

	escrever("Escrevendo", 10,34,23,6,3)
}
