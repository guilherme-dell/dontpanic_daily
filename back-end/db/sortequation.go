package db

import (
	"time"
	"math/rand"
)

func SortEquation(equation_yesterday string) (string) {
	equations := []string{
	"20*2+2",
	"10*4+2",
	"05*8+2",
	"06*8-6",
	"7*5+07",
	"51*1-9",
	"9*5-03",
	"080-38",
	"100-58",
	"32+010",
	"45-003",
	"130-88",
	"127-85",
	"66-024",
	"8+40-6",
	"6+42-6",
	"0+42-0",
	"1*44-2",
	"50-8*1",
	"70-028",
	"49-007",
	"89-047",
	"30+012",
	"022+20",
	"68/2+8",
	"88/2-2",
	"90/2-3",
	"44/1-2",
	"50/1-8",
	"40/1+2",
	"51/1-9",
}
	rand.Seed(time.Now().UnixNano())
	equation_day := equations[rand.Intn(len(equations))]
	if equation_day != equation_yesterday{
		return equation_day
	} else {
		return SortEquation(equation_day)
	}
}