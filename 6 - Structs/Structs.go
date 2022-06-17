package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco // endereco do tipo endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Jeff"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua assaí", 76}

	usuario2 := usuario{"Jeff", 22, enderecoExemplo}

	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Jeff"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 22}
	fmt.Println(usuario4)
}
