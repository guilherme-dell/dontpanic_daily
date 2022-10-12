package main

import (
	"dontpanic/db"
	"dontpanic/server"
	"fmt"
)

func main() {
	db.CreateTables()
	// db.InsertEquation(db.Equation_1)
	// db.InsertEquation(db.Equation_2)
	// db.InsertEquation(db.Equation_3)
	// db.InsertEquation(db.Equation_4)
	// db.InsertEquation(db.Equation_5)
	equation := db.GetDailyEquation()
	fmt.Println("EQUAÇÃO ESCOLHIDA:", equation)

	server.ConfiguraServidor()
}
