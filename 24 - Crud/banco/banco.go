package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL (colocar na mão)
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	conn := "root:jeff12@/banco?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", conn)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
