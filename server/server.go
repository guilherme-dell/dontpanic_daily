package server

import (
	"dontpanic/gamerules"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

var requestReceived aplicationRequest
var equacaoDoDia = [6]rune{'8', '4', '/', '2', '+', '0'}
var game gamerules.Game

func jogar(request *gin.Context) {
	body, erro := ioutil.ReadAll(request.Request.Body)
	if erro != nil {
		fmt.Println("Erro no readall")
		return
	}

	json.Unmarshal(body, &requestReceived)
	fmt.Println(requestReceived.Expressao)

	if !gamerules.TamanhoInputValido(requestReceived.Expressao){
		request.JSON(http.StatusOK,emptyField)
		fmt.Println("input invalido")
		return
	}

	input := gamerules.ConverterArrayStringParaArrayRuna(requestReceived.Expressao)
	game = gamerules.Game{input[0], input[1], input[2], input[3], input[4], input[5]}		//TODO receber um arry de strings

	if !game.ValidaEquacao(){
		request.JSON(http.StatusOK,invalidExpression)
		return
	}

	resultado := game.ResultadoExpressao()
	fmt.Println("Valor da soma: ", resultado)

	if resultado != 42{
		request.JSON(http.StatusOK, invalidResult)
		return
	}
	if game.RespostaECorreta(equacaoDoDia){
		request.JSON(http.StatusOK, gameWin(requestReceived, game))
		return
	}
	request.JSON(http.StatusOK, gameTip(requestReceived, game))
}
