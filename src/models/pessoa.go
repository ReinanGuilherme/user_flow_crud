package models

type Pessoa struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Genero         string `json:"genero"`
	DataNascimento string `json:"data_nascimento"`
	FKIDConta      int    `json:"fk_id_conta"`
}
