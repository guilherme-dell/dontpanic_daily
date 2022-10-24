package game

import (
	"errors"
	"strings"
	"unicode"
)

func validateEquation(equation string) error {
	if len([]rune(equation)) != 6 {
		return errors.New("equação com tamanho inválido")
	}
	if !isValidEquationSyntax(equation) {
		return errors.New("equação com formato inválido")
	}
	if eval(equation) != 42 {
		return errors.New("equação não resulta em 42")
	}
	return nil
}

// valida se a equação é válida
func isValidEquationSyntax(equation string) bool {
	switch {
	case !hasValidComponents(equation):
		return false
	case !hasOperator(equation):
		return false
	case hasCornersOperators(equation):
		return false
	case hasSequencialOperators(equation):
		return false
	default:
		return true
	}
}

// verifica se a equação possui ao menos um operador matemático
func hasOperator(equation string) bool {
	for _, v := range equation {
		if isOperator(v) {
			return true
		}
	}
	return false
}

// verifica se a equação infomada possui somente números e operadores matemáticos
func hasValidComponents(str string) bool {
	if len([]rune(str)) < 1 {
		return false
	}
	for _, c := range str {
		if !unicode.IsDigit(c) && !isOperator(c) {
			return false
		}
	}
	return true
}

// verifica se a equação termina ou inicia com um operador matemático
func hasCornersOperators(str string) bool {
	if (len([]rune(str))) < 2 {
		return false
	}
	for _, v := range "+-*/" {
		if strings.HasPrefix(str, string(v)) ||
			strings.HasSuffix(str, string(v)) {
			return true
		}
	}
	return false
}

func hasSequencialOperators(equation string) bool {
	lastIsOperator := false
	for _, v := range equation {
		if isOperator(v) && lastIsOperator {
			return true
		}
		lastIsOperator = isOperator(v)
	}
	return false
}

func isOperator(r rune) bool {
	switch r {
	case '*', '-', '+', '/':
		return true
	default:
		return false
	}
}
