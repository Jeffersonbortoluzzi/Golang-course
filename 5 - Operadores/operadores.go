package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 101 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	fmt.Println("----------------------------------------------------------------")
	// FIM DOS ARITMÉTICOS

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	fmt.Println("----------------------------------------------------------------")
	// FIM

	// OPERADORES RELACIONAIS
	//retornam sempre boolean
	fmt.Println(1 > 2) //retorna sempre boolean
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2) //COMPARACAO
	fmt.Println(1 != 2) //dif
	fmt.Println("----------------------------------------------------------------")
	// FIM

	// OPERADORES LÓGICOS

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("----------------------------------------------------------------")
	// FIM

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero + numero = 15
	fmt.Println(numero)

	numero--
	numero -= 3

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println("----------------------------------------------------------------")
	// FIM

	// OPERADORES TERNÁRIOS
	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
