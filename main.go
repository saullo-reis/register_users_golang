package main

import (
	"database/sql"
	"fmt"

	"registers_users/gopostgres/dbconfig"

	_ "github.com/lib/pq"
) 

func main(){
	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("conectado")
}