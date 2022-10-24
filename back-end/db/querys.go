package db

var q_CreateEquationsTable = `CREATE TABLE IF NOT EXISTS eq_day (equation CHAR(6) NOT NULL, day DATE DEFAULT CURRENT_DATE)`

var q_SelectYesterdayEquation = `SELECT equation, day FROM eq_day where day = CURRENT_DATE - 1`

var q_InsertTodayEquation = `INSERT INTO eq_day (equation) VALUES ($1)`

var q_SelectTodayEquation = `SELECT equation FROM eq_day WHERE day = CURRENT_DATE`
