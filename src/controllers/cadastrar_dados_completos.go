package controllers

import (
	"net/http"
	"time"

	"github.com/ReinanGuilherme/user_flow_crud/src/config/database"
	"github.com/ReinanGuilherme/user_flow_crud/src/repositories"
	"github.com/ReinanGuilherme/user_flow_crud/src/services"
	"github.com/ReinanGuilherme/user_flow_crud/src/utils/utils_response"
	"github.com/gin-gonic/gin"
)

type cadastrar_dados_completos_request struct {
	Pessoa repositories.Cadastrar_Pessoa_Args `json:"pessoa"`
}

func (obj *cadastrar_dados_completos_request) validar_Parametros() string {
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

	if obj.Pessoa.Data_Nascimento != "" {
		_, err := time.Parse("02-01-2006", obj.Pessoa.Data_Nascimento)
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
// @Summary .
// @Schemes
// @Description .
// @Tags Pessoa
// @Accept json
// @Produce json
// @Router /cadastrar_dados_completos [post]
// @Param Authorization header string true "Token" default(Bearer <Token aqui>)
// @Param object body cadastrar_dados_completos_request true "Objeto"
func Cadastrar_Dados_Completos(c *gin.Context) {

	var request cadastrar_dados_completos_request

	if err := c.ShouldBindJSON(&request); err != nil {
		utils_response.Response(c, http.StatusBadRequest, "Entrada no formato inválido.", nil)
		return
	}

	validar_parametros := request.validar_Parametros()
	if validar_parametros != "" {
		utils_response.Response(c, http.StatusBadRequest, validar_parametros, nil)
		return
	}

	// Iniciando a transação
	tx := database.DB.Begin()
	// Verifique se a transação foi iniciada corretamente
	if tx.Error != nil {
		utils_response.Response(c, http.StatusBadRequest, "erro ao iniciar transação", nil)
		return
	}

	retorno, err := services.Pessoa_Service(tx).Cadastrar_Pessoa(
		services.Cadastrar_Pessoa_Args{
			Pessoa: request.Pessoa,
		})
	if err != nil {
		tx.Rollback()
		utils_response.Response(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tx.Commit()
	utils_response.Response(c, http.StatusOK, "Operação realizada com sucesso.", retorno)

}
