package controllers

import (
	"net/http"

	"github.com/ReinanGuilherme/user_flow_crud/src/utils/utilsResponse"
	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Summary Rota para verificar o servidor ON.
// @Schemes
// @Description Rota para verificar o servidor ON.
// @Tags Health
// @Accept json
// @Produce json
// @Router /health [get]
func Health(c *gin.Context) {

	utilsResponse.Response(c, http.StatusOK, "Server On.", nil)
}
