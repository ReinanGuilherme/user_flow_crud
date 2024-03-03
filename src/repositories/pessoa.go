package repositories

import (
	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"gorm.io/gorm"
)

type pessoaDB struct {
	db *gorm.DB
}

func PessoaRepository(db *gorm.DB) *pessoaDB {
	return &pessoaDB{db: db}
}

type CadastrarPessoaArgs struct {
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Genero         string `json:"genero"`
	DataNascimento string `json:"data_nascimento"`
	FKIDConta      int    `json:"-"`
}

func (repository *pessoaDB) CadastrarPessoa(args CadastrarPessoaArgs) (models.Pessoa, error) {
	var retorno models.Pessoa
	err := repository.db.Raw(`
    INSERT INTO pessoas (
		nome,
		sobrenome,
		genero,
		data_nascimento,
		fk_id_conta
	)
	VALUES (?, ?, ?, ?)
	RETURNING id,nome,sobrenome,genero,data_nascimento,fk_id_conta;	
    `, args.Nome, args.Sobrenome, args.Genero, args.DataNascimento, args.FKIDConta).Scan(&retorno).Error

	if err != nil {
		return models.Pessoa{}, err
	}

	return retorno, nil
}
