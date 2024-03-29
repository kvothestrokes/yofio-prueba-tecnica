package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {

	db, err := sql.Open("mysql", "admin:12345678@tcp(database-1.cnu88ou2mbsi.us-east-2.rds.amazonaws.com:3306)/prueba_yofio")

	if err != nil {
		fmt.Printf("\nError al conectar a la base de datos: %v\n", err)
		panic(err.Error())
	}

	return db
}
