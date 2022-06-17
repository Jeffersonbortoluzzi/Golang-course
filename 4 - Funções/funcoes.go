package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculosMatematico(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 23)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	// _ para escolher qual resultado sai, sem dar erro
	_, resultadosubtracao := calculosMatematico(10, 5)
	fmt.Println(resultadosubtracao)

}
