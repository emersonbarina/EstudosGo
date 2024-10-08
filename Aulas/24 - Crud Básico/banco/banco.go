package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Drive de conexão com mysql
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "root:root@tcp(0.0.0.0:3307)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
