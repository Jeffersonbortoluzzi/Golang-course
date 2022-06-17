package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome": "Jeff",
		"sobrenome": "Bortoluzzi",
	}
	fmt.Println(usuario["nome"])

}


