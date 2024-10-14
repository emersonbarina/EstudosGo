package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar os campos de usuarios recebidos
func (usuario *Usuario) Preparar() error {
	usuario.formatar()
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o apelido é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.Trim(usuario.Nome, " ")
	usuario.Nick = strings.Trim(usuario.Nick, " ")
	usuario.Email = strings.Trim(usuario.Email, " ")
	usuario.Senha = strings.Trim(usuario.Senha, " ")
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
}
