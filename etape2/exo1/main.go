package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	var db *sql.DB
	var err error

	stringCnxn := "sqlserver://devdefis:jban4vo9@VM-DEV-2012R2-I\\SAGE2012:1433?database=comodo&encrypt=disable&parseTime=true"
	db, err = sql.Open("mssql", stringCnxn)
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Base de données ouverte")
	}

}
