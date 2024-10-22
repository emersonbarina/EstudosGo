package models

// Auth formato para atualizar a senha
type Auth struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
