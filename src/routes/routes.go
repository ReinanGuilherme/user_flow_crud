package routes

import (
	"github.com/ReinanGuilherme/user_flow_crud/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "github.com/ReinanGuilherme/user_flow_crud/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Configuração do Cors
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AddAllowHeaders("Authorization")
	r.Use(cors.New(configCors))

	//Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Configurando uma rota para verificar o servidor ON
	r.GET("/health", controllers.Health)

	//Configurando rotas da api
	SetupRouterApi(r)

	return r

}
