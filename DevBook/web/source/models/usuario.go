package models

import (
	"net/http"
	"time"
)

// Usuario representa a estrutura de uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"Nome"`
	Email       string       `json:"Email"`
	Nick        string       `json:"Nick"`
	CriadoEm    time.Time    `json:"CriadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	CanalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(CanalPublicacoes, usuarioID, r)

}

func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
