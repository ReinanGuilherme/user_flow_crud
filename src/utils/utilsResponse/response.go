package utilsResponse

import (
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// monta o objeto de resposta da requisição e retorna para o usuario a resposta.
func Response(c *gin.Context, statusCode int, message interface{}, data interface{}) {
	resp := Resp{
		Success: statusCode >= 200 && statusCode < 300,
		Message: message,
		Data:    data,
	}

	c.JSON(statusCode, resp)
}
