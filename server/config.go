package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func configurarCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"*"}, //TODO mudar os header aceitos pra somente JSON   "Content-Type", "application/json"
	}))
}

func ConfiguraServidor() {
	server := gin.Default()
	configurarCors(server)
	server.POST("/jogar", jogar)

	server.Run()
}
