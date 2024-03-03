package services

import (
	"errors"

	"github.com/ReinanGuilherme/user_flow_crud/src/models"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	"github.com/ReinanGuilherme/user_flow_crud/src/security"
	"gorm.io/gorm"
)

type contaService struct {
	db *gorm.DB
}

func ContaService(db *gorm.DB) *contaService {
	return &contaService{db: db}
}

type CadastrarContaArgs struct {
	Conta repositories.CadastrarContaArgs
}

func (service *contaService) CadastrarConta(args CadastrarContaArgs) (models.Conta, error) {

	senhaHash, err := security.ConverterSenhaParaHashBCrypt(args.Conta.Senha)
	if err != nil {
		return models.Conta{}, errors.New("erro ao tentar criptografar senha")
	}
	// alterando para o formato da senha criptografa
	args.Conta.Senha = senhaHash

	repositorio := repositories.ContaRepository(service.db)
	retorno, err := repositorio.CadastrarConta(args.Conta)
	if err != nil {
		return models.Conta{}, errors.New("erro ao tentar cadastrar conta")
	}

	return retorno, nil
}
