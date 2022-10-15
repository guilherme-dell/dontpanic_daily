package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// TodayEquation	godoc
// @Summary			Retorna a equação do dia do banco de dados após escolher uma equação aleatória
func TodayEquation() (string, error) {

	db, err := startConnection()
	if err != nil {
		return "", err
	}

	err = initializeDB(db)
	if err != nil {
		return "", err
	}
	defer db.Close()

	equation, err := selectTodayEquation(db)
	if err != nil || equation != "" {
		return equation, err
	}

	eqYesterday, err := selectYesterdayEquation(db)
	if err != nil {
		return "", err
	}

	equation = randomEquation(eqYesterday)
	err = insertTodayEquation(db, equation)
	if err != nil {
		return "", err
	}

	return equation, nil
}

// startConnection	godoc
// @Summary			Inicia a conexão com banco de dados
func startConnection() (*sql.DB, error) {

	connString := "host=db "
	connString += "port=5432 "
	connString += "user=postgres "
	connString += "password=postgres "
	connString += "dbname=dontpanic "
	connString += "sslmode=disable"

	return sql.Open("postgres", connString)
}

// initializeDB		godoc
// @Summary		Inicializa o banco de dados
func initializeDB(db *sql.DB) error {
	_, err := db.Query(q_CreateEquationsTable)
	return err
}

// selectTodayEquation		godoc
// @Summary					Seleciona a equação do dia
func selectTodayEquation(db *sql.DB) (string, error) {
	row := db.QueryRow(q_SelectTodayEquation)
	if row.Err() != nil {
		return "", row.Err()
	}

	var equation string
	row.Scan(&equation)

	return equation, nil
}

// selectYesterdayEquation	godoc
// @Summary						Seleciona a equação do dia anterior à equação do dia
func selectYesterdayEquation(db *sql.DB) (string, error) {
	equation := ""

	row := db.QueryRow(q_SelectYesterdayEquation)
	if row.Err() != nil {
		return equation, row.Err()
	}

	row.Scan(&equation)
	return equation, nil
}

// insertTodayEquation	godoc
// @Summary					Insere a equação do dia
func insertTodayEquation(db *sql.DB, equation string) error {
	_, err := db.Query(q_InsertTodayEquation, equation)
	return err
}
