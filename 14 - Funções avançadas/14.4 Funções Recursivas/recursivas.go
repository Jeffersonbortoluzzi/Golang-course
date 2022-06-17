package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	// se nao for nem 1, nem 0 tem que chamar de novo
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 4 5 8 13

	posicao := uint(10)

	for i := uint(0); i < 10; i++ {
		fmt.Println(fibonacci(posicao))

	}
	
}
