package server

import "dontpanic/gamerules"

var invalidExpression = serverResponse {
	Status : false,
	Msg_erro :"Equação invalida ",
	Expressao : RequestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

var emptyField = serverResponse {
	Status : false,
	Msg_erro :"Preencha todos os campos",
	Expressao : RequestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

var invalidResult = serverResponse {
	Status : false,
	Msg_erro :"O resultado da sua expressão não é 42",
	Expressao : RequestReceived.Expressao,
	Match : false,
	Dica : [6]string{},
}

func gameWin(r aplicationRequest, g gamerules.Game) (res serverResponse){
	res.Status  = true
	res.Msg_erro = ""
	res.Expressao = r.Expressao
	res.Match = true
	res.Dica = g.Dicas(EquacaoDoDia)

	return
}

func gameTip(r aplicationRequest, g gamerules.Game) (res serverResponse){
	res.Status  = true
	res.Msg_erro = ""
	res.Expressao = r.Expressao
	res.Match = false
	res.Dica = g.Dicas(EquacaoDoDia)

	return
}