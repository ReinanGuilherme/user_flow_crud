package repositories

import (
	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"gorm.io/gorm"
)

type pessoa_DB struct {
	db *gorm.DB
}

func Pessoa_Repository(db *gorm.DB) *pessoa_DB {
	return &pessoa_DB{db: db}
}

type Cadastrar_Pessoa_Args struct {
	Nome            string `json:"nome"`
	Sobrenome       string `json:"sobrenome"`
	Genero          string `json:"genero"`
	Data_Nascimento string `json:"data_nascimento"`
}

func (repository *pessoa_DB) Cadastrar_Pessoa(args Cadastrar_Pessoa_Args) (models.Pessoa, error) {
	var retorno models.Pessoa
	err := repository.db.Raw(`
    INSERT INTO pessoas (
		nome,
		sobrenome,
		genero,
		data_nascimento
	)
	VALUES (?, ?, ?, ?)
	RETURNING id,nome,sobrenome,genero,data_nascimento;	
    `, args.Nome, args.Sobrenome, args.Genero, args.Data_Nascimento).Scan(&retorno).Error

	if err != nil {
		return models.Pessoa{}, err
	}

	return retorno, nil
}
