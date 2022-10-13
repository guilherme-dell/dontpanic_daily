package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func Start() {
	server := gin.Default()
	configCors(server)
	server.POST("/jogar", playHandle)
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func configCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"*"}, //TODO mudar os header aceitos pra somente JSON   "Content-Type", "application/json"
	}))
}
