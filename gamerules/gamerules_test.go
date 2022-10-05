package gamerules_test

import (
	. "dontpanic/gamerules"
	"testing"
)

func TestValidaOperares1(t *testing.T) {
	game := Game{'1', '1', '+', '2', '2', '6'}

	got := game.ValidaOperadores()
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperares2(t *testing.T) {
	game := Game{'1', '1', '+', '2', '2', '+'}

	got := game.ValidaOperadores()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperares3(t *testing.T) {
	game := Game{'1', '1', '*', '/', '2', '5'}

	got := game.ValidaOperadores()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperadores4(t *testing.T) {
	game := Game{'1', '4', '5', '2', '+', '6'}
	got := game.ValidaOperadores()
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperadores5(t *testing.T) {
	game := Game{'1', '4', '5', '2', '5', '6'}
	got := game.ValidaOperadores()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperadores6(t *testing.T) {
	game := Game{'1', '+', '5', '+', '5', '6'}
	got := game.ValidaOperadores()
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaOperares_Somente_Numeros(t *testing.T) {
	game := Game{'1', '1', '2', '5', '2', '5'}

	got := game.ValidaOperadores()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaArgs1(t *testing.T) {
	game := Game{'1', '1', '+', '2', '2', '6'}
	got := game.ValidaArgs()
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaArgs2(t *testing.T) {
	game := Game{'1', '1', '+', '2', 'A', '6'}
	got := game.ValidaArgs()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaArgs3(t *testing.T) {
	game := Game{'1', '@', '+', '2', '2', '6'}
	got := game.ValidaArgs()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaArgs4(t *testing.T) {
	game := Game{}
	got := game.ValidaArgs()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaEquacao1(t *testing.T) {
	game := Game{'1', '+', '5', '+', '5', '6'}
	got := game.ValidaEquacao()
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaEquacao2(t *testing.T) {
	game := Game{'1', '+', 'g', '+', '5', '6'}
	got := game.ValidaEquacao()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestValidaEquacao3(t *testing.T) {
	game := Game{'1', '+', '*', '+', '4', '6'}
	got := game.ValidaEquacao()
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestResposta_Correta(t *testing.T) {
	correta := [6]rune{'2', '/', '1', '2', '/', '8'}
	game := Game{'2', '/', '1', '2', '/', '8'}
	got := game.RespostaECorreta(correta)
	want := true

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestResposta_Errada(t *testing.T) {
	correta := [6]rune{'2', '6', '1', '2', '/', '8'}
	game := Game{'2', '/', '1', '2', '/', '8'}
	got := game.RespostaECorreta(correta)
	want := false

	if got != want {
		t.Fatalf("Expected %t, got %t", want, got)
	}
}

func TestResultadoExpressao1(t *testing.T) {

	game := Game{'2', '0', '+', '0', '2', '2'}
	got := game.ResultadoExpressao()
	want := 42

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}

func TestResultadoExpressao2(t *testing.T) {

	game := Game{'4', '2', '-', '0', '4', '2'}
	got := game.ResultadoExpressao()
	want := 0

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}

func TestResultadoExpressao3(t *testing.T) {

	game := Game{'4', '2', '*', '1', '0', '0'}
	got := game.ResultadoExpressao()
	want := 4200

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}

func TestResultadoExpressao4(t *testing.T) {

	game := Game{'4', '2', '0', '0', '/', '2'}
	got := game.ResultadoExpressao()
	want := 2100

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}

func TestResultadoExpressao5(t *testing.T) {

	game := Game{'-', '4', '2', '-', '1', '0'}
	got := game.ResultadoExpressao()
	want := 0

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}

func TestDicas1(t *testing.T) {
	certa := [6]rune{'2', '3', '4', '2', '-', '2'}
	game := Game{'2', '3', '4', '2', '-', '2'}
	got := game.Dicas(certa)
	want := [6]string{"C", "C", "C", "C", "C", "C"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas2(t *testing.T) {
	certa := [6]rune{'0', '0', '0', '3', '+', '4'}
	game := Game{'2', '2', '5', '2', '-', '1'}
	got := game.Dicas(certa)
	want := [6]string{"X", "X", "X", "X", "X", "X"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas3(t *testing.T) {
	certa := [6]rune{'2', '3', '4', '2', '-', '2'}
	game := Game{'2', '2', '5', '2', '-', '1'}
	got := game.Dicas(certa)
	want := [6]string{"C", "T", "X", "C", "C", "X"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas4(t *testing.T) {
	certa := [6]rune{'5', '2', '7', '2', '+', '4'}
	game := Game{'2', '+', '5', '2', '7', '4'}
	got := game.Dicas(certa)
	want := [6]string{"T", "T", "T", "C", "T", "C"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas5(t *testing.T) {
	certa := [6]rune{'7', '7', '7', '+', '8', '5'}
	game := Game{'5', '*', '4', '5', '+', '4'}
	got := game.Dicas(certa)
	want := [6]string{"T", "X", "X", "T", "T", "X"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas6(t *testing.T) {
	certa := [6]rune{'7', '/', '7', '5', '7', '7'}
	game := Game{'4', '4', '8', '5', '*', '5'}
	got := game.Dicas(certa)
	want := [6]string{"X", "X", "X", "C", "X", "T"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas7(t *testing.T) {
	certa := [6]rune{'2', '/', '1', '2', '/', '8'}
	game := Game{'5', '5', '8', '/', '6', '4'}
	got := game.Dicas(certa)
	want := [6]string{"X", "X", "T", "T", "X", "X"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas8(t *testing.T) {
	certa := [6]rune{'7', '5', '*', '4', '3', '2'}
	game := Game{'7', '5', '*', '3', '1', '4'}
	got := game.Dicas(certa)
	want := [6]string{"C", "C", "C", "T", "X", "T"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas9(t *testing.T) {
	certa := [6]rune{'7', '/', '7', '5', '7', '7'}
	game := Game{'/', '5', '5', '7', '/', '5'}
	got := game.Dicas(certa)
	want := [6]string{"T", "T", "T", "T", "T", "T"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestDicas10(t *testing.T) {
	certa := [6]rune{'8', '4', '/', '2', '+', '0'}
	game := Game{'1', '0', '*', '4', '+', '2'}
	got := game.Dicas(certa)
	want := [6]string{"X", "T", "X", "T", "C", "T"}

	if CompareArrays(got, want) == false {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}