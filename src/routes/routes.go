package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server On"})
	})

	return r

}
