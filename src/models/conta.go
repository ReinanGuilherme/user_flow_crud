package models

import "time"

type Conta struct {
	ID               int       `json:"id"`
	Email            string    `json:"email"`
	Senha            string    `json:"senha"`
	Ativo            bool      `json:"ativo"`
	Token            string    `json:"token"`
	FotoPerfil       string    `json:"foto_perfil"`
	DataUltimaSessao time.Time `json:"data_ultima_sessao"`
}
