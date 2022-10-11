package db

import "fmt"

func CreateTables(){
	db := OpenConnection()
	defer db.Close()
	sql := db.QueryRow(queryCreateTable)
	if sql.Err() != nil{
		fmt.Println("Erro ai criar tabela")
	}
}