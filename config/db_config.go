package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {

	db, err := sql.Open("mysql", "root:1234@tcp(34.174.184.45:3306)/prueba_yofio")

	if err != nil {
		fmt.Printf("\nError al conectar a la base de datos: %v\n", err)
		panic(err.Error())
	}

	return db
}
