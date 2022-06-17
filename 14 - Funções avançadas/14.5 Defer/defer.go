// DEFER = ADIAR, adia a execução até o ultimo momento possivel(antes de terminar) exemplo, executa a funcao 2 antes da 1

package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

//exemplo
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O resultado será retornado") //vai ser executado antes de qualquer return.
	fmt.Println("Entrando na funcao para verificar se está aprovado")

	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(9, 8))
}
