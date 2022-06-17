package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Jeff", "Bortoluzzi", 21, 1.96}
	fmt.Println(p1)

	e1 := estudante{p1,"Engenharia", "Ufrj"}
	fmt.Println(e1)
	// fmt.Println(e1.altura) da pra puxar campos especificos.

}
