package services

import (
	"errors"

	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	utilsDate "github.com/ReinanGuilherme/user_flow_crud/src/utils/utils_date"
	"gorm.io/gorm"
)

type pessoaService struct {
	db *gorm.DB
}

func PessoaService(db *gorm.DB) *pessoaService {
	return &pessoaService{db: db}
}

type CadastrarPessoaArgs struct {
	Pessoa repositories.CadastrarPessoaArgs
}

func (service *pessoaService) CadastrarPessoa(args CadastrarPessoaArgs) (models.Pessoa, error) {

	// Formatando campo data_nascimento
	dataFormatada, err := utilsDate.FormatarDataParaPadraoPostgres(args.Pessoa.DataNascimento)
	if err != nil {
		return models.Pessoa{}, err
	}
	// Alterando o valor para a data formatado
	args.Pessoa.DataNascimento = dataFormatada

	repositorio := repositories.PessoaRepository(service.db)
	retorno, err := repositorio.CadastrarPessoa(args.Pessoa)
	if err != nil {
		return models.Pessoa{}, errors.New("erro ao tentar cadastrar pessoa")
	}

	return retorno, nil
}
