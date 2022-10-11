package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection()(*sql.DB){
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=dontpanic sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping.")
	}

	return db
}

func InsertEquation(d data_game){
	db := OpenConnection()
	defer db.Close()
	sql := db.QueryRow(queryInsertEquaion_2, d.Equation, d.Day)
	if sql.Err() != nil{
		fmt.Println("Erro inserir equação.")
	}
}

func InsertDailyEquation(equation string){
	db := OpenConnection()
	defer db.Close()
	sql := db.QueryRow(queryInsertEquation, equation)
	if sql.Err() != nil{
		fmt.Println("Erro inserir equação diaria.")
	}
}

func EquationYesterday()(equationYesterday string) {
	db := OpenConnection()
	defer db.Close()
	row := db.QueryRow(queryEquationYesterday)
	if row.Err() != nil{
		fmt.Println("Erro no select da equação de ontem.")
	}
	row.Scan(&equationYesterday)

	return
}

func VerifyEquationDaily()(bool){
	db := OpenConnection()
	defer db.Close()
	var equationDaily string
	row := db.QueryRow(queryVerifyEquationDaily)
	row.Scan(&equationDaily)
	if equationDaily == ""{
		return false
	}
	return true
}

func GetDailyEquation()(dailyEquation string){
	if !VerifyEquationDaily() {
		dailyEquation = SortEquation(EquationYesterday())
		InsertDailyEquation(dailyEquation)
		return
	}
	db := OpenConnection()
	defer db.Close()
	row := db.QueryRow(queryEquationDay)
	row.Scan(&dailyEquation)
	return
}
