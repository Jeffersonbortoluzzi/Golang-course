package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução...")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 - n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 7")
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 7))
	fmt.Println("Pós execucao")
}
