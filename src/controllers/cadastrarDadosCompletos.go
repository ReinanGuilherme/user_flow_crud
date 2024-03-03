package controllers

import (
	"net/http"
	"time"

	"github.com/ReinanGuilherme/user_flow_crud/src/config/database"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	"github.com/ReinanGuilherme/user_flow_crud/src/services"
	"github.com/ReinanGuilherme/user_flow_crud/src/utils/utilsResponse"
	"github.com/gin-gonic/gin"
)

type cadastrarDadosCompletosRequest struct {
	Conta  repositories.CadastrarContaArgs  `json:"conta"`
	Pessoa repositories.CadastrarPessoaArgs `json:"pessoa"`
}

func (obj *cadastrarDadosCompletosRequest) validarParametros() string {
	var erros string = ""
	if obj.Pessoa.Nome == "" {
		erros += "nome inválido,"
	}

	if obj.Pessoa.Sobrenome == "" {
		erros += "sobrenome inválido,"
	}

	if obj.Pessoa.Genero == "" {
		erros += "genero inválido,"
	}

	if obj.Pessoa.DataNascimento != "" {
		_, err := time.Parse("02-01-2006", obj.Pessoa.DataNascimento)
		if err != nil {
			erros += "data_nascimento inválida,"
		}

	}

	//removendo o ultimo caracter (,) da string
	if len(erros) > 0 {
		erros = "Erro - " + erros[:len(erros)-1]
	}

	return erros
}

// @BasePath /
// @Summary Cadastrar uma nova conta no sistema com seus dados completos como pessoa, endereço e contato.
// @Schemes
// @Description Cadastrar uma nova conta no sistema com seus dados completos como pessoa, endereço e contato.
// @Tags Cadastrar
// @Accept json
// @Produce json
// @Router /cadastrarDadosCompletos [post]
// @Param object body cadastrarDadosCompletosRequest true "Objeto"
func CadastrarDadosCompletos(c *gin.Context) {

	var request cadastrarDadosCompletosRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.Response(c, http.StatusBadRequest, "Entrada no formato inválido.", nil)
		return
	}

	validar_parametros := request.validarParametros()
	if validar_parametros != "" {
		utilsResponse.Response(c, http.StatusBadRequest, validar_parametros, nil)
		return
	}

	// Iniciando a transação
	tx := database.DB.Begin()
	// Verifique se a transação foi iniciada corretamente
	if tx.Error != nil {
		utilsResponse.Response(c, http.StatusBadRequest, "erro ao iniciar transação", nil)
		return
	}

	retornoConta, err := services.ContaService(tx).CadastrarConta(
		services.CadastrarContaArgs{
			Conta: request.Conta,
		})
	if err != nil {
		tx.Rollback()
		utilsResponse.Response(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Adicionando o id da conta criada para gerar o relacionamento com a tabela pessoas
	request.Pessoa.FKIDConta = retornoConta.ID
	retornoPessoa, err := services.PessoaService(tx).CadastrarPessoa(
		services.CadastrarPessoaArgs{
			Pessoa: request.Pessoa,
		})
	if err != nil {
		tx.Rollback()
		utilsResponse.Response(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tx.Commit()
	utilsResponse.Response(c, http.StatusCreated, "Operação realizada com sucesso.", gin.H{
		"conta":  retornoConta,
		"pessoa": retornoPessoa,
	})

}
