package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia é diferente de Paralelismo!
	go escrever("Olá mundo") //"go" nao espera o comando anterior acabar para executar
	escrever("Programa em Golang!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
