package server

import (
	"dontpanic/gamerules"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type serverResponse struct {
	Status		bool		`json:"status"`
	Msg_erro	string		`json:"msg_erro"`
	Expressao	[6]string	`json:"expressao"`
	Match		bool		`json:"match"`
	Dica		[6]string	`json:"dica"`
}

type aplicationRequest struct {
	Expressao [6]string `json:"expressao"`
}

var RequestReceived aplicationRequest
var Response serverResponse
var EquacaoDoDia = [6]rune{'8', '4', '/', '2', '+', '0'}
var game gamerules.Game

func jogar(request *gin.Context) {
	body, erro := ioutil.ReadAll(request.Request.Body)
	if erro != nil {
		fmt.Println("Erro no readall")
	}

	json.Unmarshal(body, &RequestReceived)
	fmt.Println(RequestReceived.Expressao)

	if !gamerules.TamanhoInputValido(RequestReceived.Expressao){
		request.JSON(http.StatusOK,emptyField)
		fmt.Println("input invalido")
		return
	}

	input := gamerules.ConverterArrayStringParaArrayRuna(RequestReceived.Expressao)
	game = gamerules.Game{input[0], input[1], input[2], input[3], input[4], input[5]}

	if !game.ValidaEquacao(){
		request.JSON(http.StatusOK,invalidExpression)
		return
	}

	resultado := game.ResultadoExpressao()
	fmt.Println("Valor da soma: ", resultado)

	if resultado != 42{
		request.JSON(200, invalidResult)
		return
	}
	if game.RespostaECorreta(EquacaoDoDia){

		request.JSON(200, gameWin(RequestReceived, game))
		return
	}
	if !game.RespostaECorreta(EquacaoDoDia){
		fmt.Println("input:", input)
		request.JSON(200, gameTip(RequestReceived, game))
		return
	}
}

func configurarCors(server *gin.Engine){
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"*"},
	}))
}

func ConfiguraServidor() {
	server := gin.Default()
	configurarCors(server)
	server.POST("/jogar", jogar)

	server.Run()
}
