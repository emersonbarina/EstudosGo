package controllers

import (
	"api/source/auth"
	"api/source/banco"
	"api/source/models"
	"api/source/repositories"
	"api/source/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao adiciona um nova publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao

	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes traz publicações do feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {}

// AtualizarPublicacao altera os dados da publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}

// DeletarPublicacao deleta uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {}
