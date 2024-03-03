package routes

import (
	"github.com/ReinanGuilherme/user_flow_crud/src/controllers"
	"github.com/gin-gonic/gin"
)

// configurações das rotas da API
func SetupRouterApi(r *gin.Engine) {

	//publica
	rotaPublica := r.Group("/")
	rotaPublica.POST("/cadastrarConta", controllers.CadastrarConta)
	rotaPublica.POST("/cadastrarDadosCompletos", controllers.CadastrarDadosCompletos)

}
