package game

import (
	"strconv"
	"strings"
)

type Game interface {
	Hints(string) (string, error)
}

// tipo que contém a equação oculta do jogo
type game string

// cria um novo jogo com a equação oculta
func New(hiddenEquation string) (game, error) {
	err := validateEquation(hiddenEquation)
	if err != nil {
		return "", err
	}
	return game(hiddenEquation), nil
}

// representação do jogo como estring
func (g game) String() string {
	return string(g)
}

// informa as dicas da equação com base na equação informada
func (g game) Hints(equation string) (string, error) {

	err := validateEquation(equation)
	if err != nil {
		return "", err
	}

	gameRunes := []rune(g)
	eqRunes := []rune(equation)
	hints := []rune{'#', '#', '#', '#', '#', '#'}
	incorrectPos := make(map[rune]int)

	for i, v := range gameRunes {
		if v == eqRunes[i] {
			hints[i] = 'C'
			continue
		}
		incorrectPos[v]++
	}

	for i, v := range hints {
		if v != '#' {
			continue
		}

		if incorrectPos[eqRunes[i]] > 0 {
			hints[i] = 'T'
			incorrectPos[v]--
			continue
		}
		hints[i] = 'X'
	}

	return string(hints), nil
}

// avalia a equação informada a calcula o seu valor
func eval(equation string) int {
	values := values(equation)
	operators := operators(equation)

	result := 0
	if len(values) == 3 && len(operators) == 2 {
		if operators[1] == "*" || operators[1] == "/" {
			result = calculate(values[1], values[2], operators[1])
			result = calculate(values[0], result, operators[0])
		} else {
			result = calculate(values[0], values[1], operators[0])
			result = calculate(result, values[2], operators[1])
		}
	}

	if len(values) == 2 && len(operators) == 1 {
		result = calculate(values[0], values[1], operators[0])
	}

	return result
}

// retorna uma slice de int contendo somente os números da equação
func values(equation string) []int {
	var values []int
	onlyValues := strings.FieldsFunc(equation, isOperator)
	for _, v := range onlyValues {
		intValue, _ := strconv.Atoi(v)
		values = append(values, intValue)
	}
	return values
}

// retorna uma slice de string contendo somente os operadores da equação
func operators(equation string) []string {
	var operators []string
	for _, v := range equation {
		if isOperator(v) {
			operators = append(operators, string(v))
		}
	}
	return operators
}

// calcula dois valores utilizando a operação informada
func calculate(val1, val2 int, operation string) int {
	switch operation {
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	case "-":
		return val1 - val2
	case "+":
		return val1 + val2
	default:
		return 0
	}
}
