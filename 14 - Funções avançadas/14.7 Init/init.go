package main

import "fmt"

var n int

func main() {
	fmt.Println("Executando a func main")
	fmt.Println(n)
}

//setar valor ou algo do tipo.
func init() {
	fmt.Println("executando a func init (executa primeiro)")
	n = 10
}
