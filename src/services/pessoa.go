package services

import (
	"errors"

	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	"github.com/ReinanGuilherme/user_flow_crud/src/utils/utils_date"
	"gorm.io/gorm"
)

type pessoa_service struct {
	db *gorm.DB
}

func Pessoa_Service(db *gorm.DB) *pessoa_service {
	return &pessoa_service{db: db}
}

type Cadastrar_Pessoa_Args struct {
	Pessoa repositories.Cadastrar_Pessoa_Args
}

func (service *pessoa_service) Cadastrar_Pessoa(args Cadastrar_Pessoa_Args) (models.Pessoa, error) {

	// Formatando campo data_nascimento
	data_formatada, err := utils_date.Formatar_Data_Para_Padrao_Postgres(args.Pessoa.Data_Nascimento)
	if err != nil {
		return models.Pessoa{}, err
	}
	// Alterando o valor para a data formatado
	args.Pessoa.Data_Nascimento = data_formatada

	repositorio := repositories.Pessoa_Repository(service.db)
	retorno, err := repositorio.Cadastrar_Pessoa(args.Pessoa)
	if err != nil {
		return models.Pessoa{}, errors.New("erro ao tentar cadastrar pessoa")
	}

	return retorno, nil
}
