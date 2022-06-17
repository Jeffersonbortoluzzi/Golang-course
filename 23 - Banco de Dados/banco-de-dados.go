package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn := "root:jeff12@tcp(127.0.0.1:3306)/Banco?charset=utf8&parsetime=True&loc=Local"
	db, erro := sql.Open("mysql", conn)
	if erro != nil {
		fmt.Println("Erro dentro do SQL open")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro dentro do ping")
		log.Fatal(erro)
	}
	fmt.Println("Connexão bem sucedida!")

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		fmt.Println("Erro dentro do query")
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
