package server

import "dontpanic/gamerules"

var invalidExpression = serverResponse {
	Status : false,
	Msg_erro :"Equação invalida ",
	Expressao : requestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

var emptyField = serverResponse {
	Status : false,
	Msg_erro :"Preencha todos os campos",
	Expressao : requestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

var invalidResult = serverResponse {
	Status : false,
	Msg_erro :"O resultado da sua expressão não é 42",
	Expressao : requestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

func gameWin(r aplicationRequest, g gamerules.Game) (response serverResponse){
	response.Status  = true
	response.Msg_erro = ""
	response.Expressao = r.Expressao
	response.Match = true
	response.Dica = g.Dicas(equacaoDoDia)

	return
}

func gameTip(r aplicationRequest, g gamerules.Game) (response serverResponse){
	response.Status  = true
	response.Msg_erro = ""
	response.Expressao = r.Expressao
	response.Match = false
	response.Dica = g.Dicas(equacaoDoDia)

	return
}