package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Bob", "Golden", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)
	// newBuffer: ao chamar, dá uma saida de bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	//sem usar struct:

	c2 := map[string]string{
		"nome": "Jerry",
		"raca": "Poodle",
	}
	cachorro2EmJason, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJason)
}
