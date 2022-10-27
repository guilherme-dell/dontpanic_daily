function fazPost(url, body){
	let request = new XMLHttpRequest()
	request.open("POST", url, true)
	request.setRequestHeader("Content-Type", "application/json")
	request.send(JSON.stringify(body))
	console.log("REQUEST ENVIADO = ", body)
	request.onload = function (){
		console.log("Resposta api:", this.responseText)
		respose_api = JSON.parse(this.responseText)
		printa_dica(respose_api)
	}
	return request.responseText
}

function printa_dica(response){

	limpar_historico_respostas()
	limpar_dicas()
	limpar_msg_erro()
	limpar_msg_vitoria()

	for (i = 0; i < 6; i++)
	{
		pos = i + 1
		campo_historico = document.getElementById("input-" + pos)

		if (response.dica[i] == "C") {
			inserir_historico_resposta(pos, campo_historico.value, true)
		}
		else {
			inserir_historico_resposta(pos, campo_historico.value)
		}

		inserir_dica(pos, response.dica[i])
	}

	if (response.status == false && response.msg_erro.length > 0)
		inserir_msg_erro(response.msg_erro)

	if (response.status && response.msg_erro.lenght > 0)
		inserir_msg_erro(response.msg_erro)

	if (response.status && response.match)
		inserir_msg_vitoria("Você Ganhou!")

	limpar_resposta()
	console.log(response)
}

function jogar(){
	event.preventDefault()
	console.log("request enviado")
	let url = "/jogar"
	let input1 = document.getElementById("input-1").value
	let input2 = document.getElementById("input-2").value
	let input3 = document.getElementById("input-3").value
	let input4 = document.getElementById("input-4").value
	let input5 = document.getElementById("input-5").value
	let input6 = document.getElementById("input-6").value

	body = {
		"expressao":[input1, input2, input3, input4, input5, input6]
	}
	resposta_post = fazPost(url, body)
}

function limpar_dicas() {
	for (i = 1; i <= 6; i++) {
		element = document.getElementById("campo_dica" + i)
		element.classList.remove("tip-c", "tip-x", "tip-t")
		element.innerHTML = ""
	}
}

function limpar_dica(n) {
	element = document.getElementById("campo_dica" + n)
	element.classList.remove("tip-c", "tip-x", "tip-t")
	element.innerHTML = ""
}

function inserir_dica(n, tipo) {
	limpar_dica(n)
	element = document.getElementById("campo_dica" + n)
	if (tipo == "T") element.classList.add("tip-t")
	if (tipo == "C") element.classList.add("tip-c")
	if (tipo == "X") element.classList.add("tip-x")
	element.innerHTML = tipo
}

function limpar_historico_respostas() {
	for (i = 1; i <= 6; i++) {
		element = document.getElementById("last_input" + i)
		element.classList.remove("last-answer", "light")
		element.innerHTML = ""
	}
}

function inserir_historico_resposta(n, resp, light = false) {
	element = document.getElementById("last_input" + n)
	element.classList.add("last-answer")
	if (light) element.classList.add("light")
	element.innerHTML = resp
}

function limpar_resposta() {
	for (i = 1; i <= 6; i++) {
		element = document.getElementById("input-" + i)
		element.value = ""
	}
}

function limpar_msg_erro() {
	element = document.getElementById("error-message")
	element.classList.remove("show-error")
}

function inserir_msg_erro(msg) {
	element = document.getElementById("error-message")
	element.classList.add("show-error")
	element.innerHTML = msg
}

function limpar_msg_vitoria() {
	element = document.getElementById("win-message")
	element.classList.remove("show-message")
}

function inserir_msg_vitoria(msg) {
	element = document.getElementById("win-message")
	element.classList.add("show-message")
	element.innerHTML = msg
}

function limpar_game() {
	limpar_dicas()
	limpar_historico_respostas()
	limpar_resposta()
	ocultar_loading()
	limpar_msg_erro()
}
