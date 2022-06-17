package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	//enquanto J menor que 10, será incrementado um por um
	for j := 0; j < 10; j += 5 { //incrementa de 5 em 5 {
		fmt.Println("Incrementando j", j) //,j está sendo printado
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Jeff", "John", "Pedro"}

	for _, nomes := range nomes {
		fmt.Println(nomes)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Henrique",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}
