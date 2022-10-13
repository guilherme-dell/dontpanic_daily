package server

// configura e retorna uma resposta de erro
func errorResponse(equation [6]string, msg string) serverResponse {
	return serverResponse{
		Status:   false,
		Msg_erro: msg,
		Equation: equation,
		Match:    false,
		Hints:    [6]string{},
	}
}

// configura e retorna uma resposta de sucesso
func successResponse(equation [6]string, match bool, hints [6]string) serverResponse {
	return serverResponse{
		Status:   false,
		Msg_erro: "",
		Equation: equation,
		Match:    match,
		Hints:    hints,
	}
}
