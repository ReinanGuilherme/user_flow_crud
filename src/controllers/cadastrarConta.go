package controllers

import (
	"net/http"

	"github.com/ReinanGuilherme/user_flow_crud/src/config/database"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	"github.com/ReinanGuilherme/user_flow_crud/src/services"
	"github.com/ReinanGuilherme/user_flow_crud/src/utils/utilsResponse"
	"github.com/gin-gonic/gin"
)

type cadastrarContaRequest struct {
	Conta repositories.CadastrarContaArgs `json:"conta"`
}

func (obj *cadastrarContaRequest) validarParametros() string {
	var erros string = ""
	if obj.Conta.Email == "" {
		erros += "email inválido,"
	}

	//removendo o ultimo caracter (,) da string
	if len(erros) > 0 {
		erros = "Erro - " + erros[:len(erros)-1]
	}

	return erros
}

// @BasePath /
// @Summary Cadastrar uma nova conta no sistema.
// @Schemes
// @Description Cadastrar uma nova conta no sistema.
// @Tags Cadastrar
// @Accept json
// @Produce json
// @Router /cadastrarConta [post]
// @Param object body cadastrarContaRequest true "Objeto"
func CadastrarConta(c *gin.Context) {

	var request cadastrarContaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utilsResponse.Response(c, http.StatusBadRequest, "Entrada no formato inválido.", nil)
		return
	}

	validar_parametros := request.validarParametros()
	if validar_parametros != "" {
		utilsResponse.Response(c, http.StatusBadRequest, validar_parametros, nil)
		return
	}

	retorno, err := services.ContaService(database.DB).CadastrarConta(
		services.CadastrarContaArgs{
			Conta: request.Conta,
		})
	if err != nil {
		utilsResponse.Response(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utilsResponse.Response(c, http.StatusCreated, "Operação realizada com sucesso.", retorno)

}
