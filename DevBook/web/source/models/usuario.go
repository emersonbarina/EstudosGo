package models

import "time"

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
