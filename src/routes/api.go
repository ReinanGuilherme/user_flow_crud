package routes

import (
	"github.com/ReinanGuilherme/user_flow_crud/src/controllers"
	"github.com/gin-gonic/gin"
)

// verifica se o servidor está funcionando ou não.
func SetupRouterApi(r *gin.Engine) {

	//publica
	rota_publica := r.Group("/")
	rota_publica.POST("/cadastrar_pessoa", controllers.Cadastrar_Pessoa)

}
