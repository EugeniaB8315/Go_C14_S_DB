package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dataSource := "user:pass@tcp(server:Port)/dbName"
	// Open inicia un pool de conexiones. Solo abrir una vez
	storageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
}
