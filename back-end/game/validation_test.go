package game

import (
	"testing"
)

func Test_hasOperator(t *testing.T) {

	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", false},
		{"+", true},
		{"-", true},
		{"/", true},
		{"*", true},
		{"6", false},
		{"9587", false},
		{"+/-*", true},
		{"5+-9*0", true},
		{"a", false},
		{"p=ag091$#^", false},
		{"as+iure", true},
		{"59+456", true},
		{"69-23", true},
		{"41/23", true},
		{"4*214", true},
		{"+4*214", true},
		{"41/23-", true},
		{"+69-23**", true},
	}
	for _, test := range cases {
		got := hasOperator(test.input)
		if got != test.want {
			t.Fatalf("Expected %t, got %t", test.want, got)
		}
	}
}

func Test_hasValidComponents(t *testing.T) {

	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", false},
		{"+", true},
		{"-", true},
		{"/", true},
		{"*", true},
		{"6", true},
		{"9587", true},
		{"+/-*", true},
		{"5+-9*0", true},
		{"a", false},
		{"as#4%", false},
		{"a6sd9e", false},
	}
	for _, test := range cases {
		got := hasValidComponents(test.input)
		if got != test.want {
			t.Fatalf("Expected %t, got %t", test.want, got)
		}
	}
}

func Test_hasCornersOperators(t *testing.T) {

	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", false},
		{"+", false},
		{"++", true},
		{"+-", true},
		{"/*", true},
		{"-5", true},
		{"5*", true},
		{"+595", true},
		{"959*", true},
		{"-69+", true},
	}

	for _, test := range cases {
		got := hasCornersOperators(test.input)
		if got != test.want {
			t.Fatalf("Expected %t, got %t", test.want, got)
		}
	}
}

func Test_hasSequentialOperation(t *testing.T) {

	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{" ", false},
		{"+", false},
		{"-", false},
		{"/", false},
		{"*", false},
		{"++", true},
		{"//", true},
		{"**", true},
		{"--", true},
		{"-+", true},
		{"/*", true},
		{"+/*-", true},
		{"+5-8+-9", true},
		{"569878", false},
		{"-569/878", false},
	}

	for _, test := range cases {
		got := hasSequencialOperators(test.input)
		if got != test.want {
			t.Fatalf("Expected %t, got %t", test.want, got)
		}
	}
}

func Test_isOperator(t *testing.T) {

	cases := []struct {
		input rune
		want  bool
	}{
		{' ', false},
		{'*', true},
		{'/', true},
		{'-', true},
		{'+', true},
		{'^', false},
		{'a', false},
		{'9', false},
	}

	for _, test := range cases {
		got := isOperator(test.input)
		if got != test.want {
			t.Fatalf("Expected %t, got %t", test.want, got)
		}
	}
}
