package models

type Pessoa struct {
	ID              int    `json:"id"`
	Nome            string `json:"nome"`
	Sobrenome       string `json:"sobrenome"`
	Genero          string `json:"genero"`
	Data_Nascimento string `json:"data_nascimento"`
}
