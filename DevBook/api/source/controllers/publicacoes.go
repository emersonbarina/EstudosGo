package controllers

import "net/http"

// CriarPublicacao adiciona um nova publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacoes traz publicações do feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {}

// AtualizarPublicacao altera os dados da publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}

// DeletarPublicacao deleta uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {}
