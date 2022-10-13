package server

import (
	"dontpanic/db"
	"dontpanic/game"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func playHandle(gc *gin.Context) {

	// carrega o corpo da requisição
	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var clientReq clientRequest
	json.Unmarshal(body, &clientReq)

	// request com campos em branco
	if hasEmptyString(clientReq.Equation[:]) {
		gc.JSON(http.StatusOK, errorResponse(
			clientReq.Equation, "preencha todos os campos"))
		return
	}

	// pega do bando a equação do dia
	todayEquation, err := db.TodayEquation()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var dontpanic game.Game

	// inicializa o jogo com a equação do dia
	dontpanic, err = game.New(todayEquation)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// pega as dicas e caso haja alguma divergência na equação
	// envia a mensagem de erro para o client
	hints, err := dontpanic.Hints(sliceToString(clientReq.Equation[:]))
	if err != nil {
		gc.JSON(http.StatusOK, errorResponse(clientReq.Equation, err.Error()))
		return
	}

	// envia as dicas para o client
	match := hints == "CCCCCC"
	gc.JSON(http.StatusOK, successResponse(clientReq.Equation, match, hintsAdapter(hints)))
}

// transforma a string de dica em um array de strings
func hintsAdapter(hints string) [6]string {
	res := [6]string{}
	for i, v := range hints {
		if i >= 6 {
			break
		}
		res[i] = string(v)
	}
	return res
}

// transforma um slice em uma string
func sliceToString(slice []string) string {
	result := ""
	for _, v := range slice {
		result += v
	}
	return result
}

// verifica se há strings vazias na slice informada
func hasEmptyString(slice []string) bool {
	if len(slice) < 1 {
		return true
	}
	for _, v := range slice {
		if strings.TrimSpace(v) == "" {
			return true
		}
	}
	return false
}
