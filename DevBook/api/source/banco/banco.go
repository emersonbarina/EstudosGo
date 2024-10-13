package banco

import (
	"api/source/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Drive de conexão com mysql
)

// Conectar - Abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	// se houve algum problema, a conexão será fechada
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
