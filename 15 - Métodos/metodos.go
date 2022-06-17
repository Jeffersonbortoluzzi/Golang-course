package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//METODO, o "u" pode ser qualquer coisa, mas o 2 tem que ser o mesmo nome do tipo
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco!", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{nome: "Usuario 1", idade: 17}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{nome: "Jeff ", idade: 22}
	fmt.Println(usuario2)
	usuario1.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
