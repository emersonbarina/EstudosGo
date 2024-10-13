package repositories

import (
	"api/source/models"
	"database/sql"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um nova instância de repositório usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário do banco de dados
func (u Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	return 0, nil
}
