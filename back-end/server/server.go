package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "dontpanic/docs"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

type serverResponse struct {
	Status   bool      `json:"status"`
	Msg_erro string    `json:"msg_erro"`
	Equation [6]string `json:"expressao"`
	Match    bool      `json:"match"`
	Hints    [6]string `json:"dica"`
}

type clientRequest struct {
	Equation [6]string `json:"expressao"`
}

// Start		godoc
// @Summary Inicializa a API com GIN
func Start() {
	server := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	configCors(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.POST("/jogar", playHandle)
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// configCors		godoc
// @Summary		Configuração de Cors
func configCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"*"}, //TODO mudar os header aceitos pra somente JSON   "Content-Type", "application/json"
	}))
}
