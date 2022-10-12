package gamerules

import (
	"fmt"
	"strconv"
	"unicode"
)

type Game [6]rune

// Retorna true se não houver operadores nas extremidades
// e se não houver dois em seguida
func (g Game) ValidaOperadores() bool {

	// Verifica operador nas extremidades
	if isOperator(g[0]) || isOperator(g[5]) {
		return false
	}

	// Verifica se nao ha operadores consecutivos
	for i := 0; i < len(g); i++ {
		if (isOperator(g[i])) && i != len(g)-1 {
			if isOperator(g[i+1]) {
				return false
			}
		}
	}

	// Verifica se ha pelo menos um operador
	count := 0
	for _, arg := range g {
		if isOperator(arg) {
			count++
		}
	}

	return count > 0
}

func (g Game) ValidaArgs() bool {
	for _, arg := range g {
		if !unicode.IsDigit(arg) && !isOperator(arg) {
			return false
		}
	}
	return true
}

func (g Game) ValidaEquacao() bool {
	if g.ValidaArgs() && g.ValidaOperadores() {
		return true
	}
	return false
}

// Válida se a equação tentada é igual a correta
func (g Game) RespostaECorreta(correta [6]rune) bool {
	for i := range correta {
		if g[i] != correta[i] {
			return false
		}
	}
	return true
}

// converte array de string em números. PRECISA DESSA FUNÇÃO?
func RetornaArry(argv [6]string) [6]int {

	array_retorno := [6]int{0, 0, 0, 0, 0, 0}

	for i := 0; i < len(argv); i++ {
		convert, err := strconv.Atoi(argv[i])

		if err != nil {
			fmt.Println("Eroo no Atoi")
			fmt.Println(err)
		}
		array_retorno[i] = convert
	}
	return array_retorno

}

func isOperator(o rune) bool {
	switch o {
	case '*', '-', '+', '/':
		return true
	default:
		return false
	}
}

func CompareArrays(a, b [6]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (g Game) Dicas(certa [6]rune) [6]string {
	var dica = make([]string, 0)
	var dicaArray [6]string

	count := 0
	fix := g[0]
	for i := 0; i < 6; i++ {
		if fix == certa[i] {
			dica = append(dica, "C")
		}
		for j := 1; j < 6; j++ {
			if (fix == certa[0] && certa[0] != certa[i]) ||
				fix == certa[j] && fix != certa[i] {
				dica = append(dica, "T")
				break
			} else {
				count = count + 1
			}
			if count == 5 && fix != certa[i] {
				dica = append(dica, "X")
			}
		}
		count = 0
		if i != 5 {
			fix = g[i+1]
		}
	}
	copy(dicaArray[:], dica[:6])
	return dicaArray
}

func (g Game) toStringArray() [6]string {
	var str [6]string

	for i, v := range g {
		str[i] = string(v)
	}

	return str
}

func (g Game) ResultadoExpressao() int {
	if g.ValidaEquacao() {
		return somar(g.toStringArray())
	}
	return 0
}

func isStringOperator(o string) bool {
	switch o {
	case "*", "-", "+", "/":
		return true
	default:
		return false
	}
}

func calculate(a, b int, op string) int {
	switch op {
	case "*":
		return a * b
	case "/":
		return a / b
	case "-":
		return a - b
	case "+":
		return a + b
	default:
		return 0
	}
}

func ConcatenaNumero(argumentos [6]string, start int) (total, posi int) {

	var amontuado string
	i := start
	for ; i <= 5 && !isStringOperator(argumentos[i]); i++ {
		amontuado += argumentos[i]
	}
	total, _ = strconv.Atoi(amontuado)
	posi = i
	return total, posi
}

func TresValores(argumentos [6]string) bool {
	counter := 0
	for i := 0; i < len(argumentos); i++ {
		if isStringOperator(argumentos[i]) {
			counter++
		}
	}
	return counter == 2
}

func somar(argumentos [6]string) int {
	if TresValores(argumentos) {
		valor1, operador_pos := ConcatenaNumero(argumentos, 0)
		operador1 := argumentos[operador_pos]
		valor2, operador_pos2 := ConcatenaNumero(argumentos, operador_pos+1)
		operador2 := argumentos[operador_pos2]
		valor3, _ := ConcatenaNumero(argumentos, operador_pos2+1)

		var resultado int
		if operador2 == "*" || operador2 == "/" {
			resultado = calculate(valor2, valor3, operador2)
			resultado = calculate(valor1, resultado, operador1)
		} else {
			resultado = calculate(valor1, valor2, operador1)
			resultado = calculate(resultado, valor3, operador2)
		}
		return resultado

	} else {
		valor1, operador_pos := ConcatenaNumero(argumentos, 0)
		operador1 := argumentos[operador_pos]
		valor2, _ := ConcatenaNumero(argumentos, operador_pos+1)
		resultado := calculate(valor1, valor2, operador1)
		return resultado
	}
}

func TamanhoInputValido(input [6]string) bool {
	for _, valor := range input {
		if len(valor) != 1 {
			return false
		}
	}
	return true
}

func ConverterArrayStringParaArrayRuna(input [6]string) [6]rune {
	var arrayDeRuna [6]rune
	for i, valor := range input {
		arrayDeRuna[i] = rune(valor[0])
	}

	return arrayDeRuna
}

func ConvertsorRuna(input string) [6]rune {
	var arrayDeRuna [6]rune
	for i:=0; i <= 5; i++ {
		arrayDeRuna[i] = rune(input[i])
	}

	return arrayDeRuna
}