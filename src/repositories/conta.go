package repositories

import (
	"time"

	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"gorm.io/gorm"
)

type contaDB struct {
	db *gorm.DB
}

func ContaRepository(db *gorm.DB) *contaDB {
	return &contaDB{db: db}
}

type CadastrarContaArgs struct {
	Email            string    `json:"email"`
	Senha            string    `json:"senha"`
	Ativo            bool      `json:"ativo"`
	Token            string    `json:"-"`
	FotoPerfil       string    `json:"foto_perfil"`
	DataUltimaSessao time.Time `json:"-"`
}

func (repository *contaDB) CadastrarConta(args CadastrarContaArgs) (models.Conta, error) {
	var retorno models.Conta
	err := repository.db.Raw(`
    INSERT INTO contas (
		email,
		senha,
		ativo,
		token,
		foto_perfil,
		data_ultima_sessao
	)
	VALUES (?, ?, ?, null, ?, null)
	RETURNING id, email, senha, ativo, token, foto_perfil, data_ultima_sessao;	
    `, args.Email, args.Senha, args.Ativo, args.FotoPerfil).Scan(&retorno).Error

	if err != nil {
		return models.Conta{}, err
	}

	return retorno, nil
}
