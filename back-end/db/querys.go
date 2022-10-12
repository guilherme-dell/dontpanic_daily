package db

var queryCreateTable = `CREATE TABLE IF NOT EXISTS eq_day (equation CHAR(6) NOT NULL, day DATE DEFAULT CURRENT_DATE)`

var queryEquationYesterday = `SELECT equation, day FROM eq_day where day = CURRENT_DATE - 1`

var queryInsertEquation = `INSERT INTO eq_day (equation) VALUES ($1)`

var queryInsertEquaion_2 = `INSERT INTO eq_day (equation,day) VALUES ($1,$2)`

var queryVerifyEquationDaily = `SELECT day FROM eq_day where day = CURRENT_DATE`

var queryEquationDay = `SELECT equation FROM eq_day WHERE day = CURRENT_DATE`

type data_game struct {
	Equation	string
	Day			string
}

var Equation_1 = data_game{
	Equation: "84/2-0",
	Day: "2022-10-06",
}
var Equation_2 = data_game{
	Equation: "84/2+0",
	Day: "2022-10-07",
}
var Equation_3 = data_game{
	Equation: "127-85",
	Day: "2022-10-08",
}
var Equation_4 = data_game{
	Equation: "66-024",
	Day: "2022-10-09",
}
var Equation_5 = data_game{
	Equation: "84/2+0",
	Day: "2022-10-10",
}
