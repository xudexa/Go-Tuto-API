package main

/*
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [contact](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[nom] [nvarchar](255) NOT NULL,
	[Prenom] [nvarchar](255) NULL,
	PRIMARY KEY (ID)
) ON [PRIMARY]

GO
*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type databaseSetting struct {
	SettingName  string `json:"name"`
	DatabaseName string `json:"dbName"`
	HostName     string `json:"server"`
	Port         int    `json:"port"`
	Driver       string `json:"driver"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

func openDB() (*sql.DB, error) {
	var dbSetting databaseSetting

	dbSetting.SettingName = "MaConfig"
	dbSetting.Driver = "mssql"
	dbSetting.HostName = "localhost"
	dbSetting.Port = 1433
	dbSetting.User = "devdefis"
	dbSetting.Password = "jban4vo9"
	dbSetting.DatabaseName = "defistuto"

	stringCnxn := "sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable&parseTime=true"
	stringCnxn = fmt.Sprintf(stringCnxn,
		dbSetting.User,
		dbSetting.Password,
		dbSetting.HostName,
		dbSetting.Port,
		dbSetting.DatabaseName)

	db, err := sql.Open(dbSetting.Driver, stringCnxn)

	return db, err
}

func main() {

	var db *sql.DB
	var err error

	db, err = openDB()
	defer db.Close()
	if err == nil {
		var nom string
		var prenom string
		var id int

		rows, errQuery := db.Query("Select * from contact")
		if errQuery == nil {
			defer rows.Close()
			for rows.Next() {

				errScan := rows.Scan(&id, &nom, &prenom)
				if errScan == nil {
					fmt.Println(fmt.Sprintf("ID : %d, Nom : %s, Prénom: %s", id, nom, prenom))
				} else {
					err = errScan
				}
			}
		} else {
			err = errQuery
		}

		fmt.Printf("Lecture du contact 1 \n")

		row := db.QueryRow("SELECT * FROM contact WHERE id = $1;", 1)
		errScan := row.Scan(&id, &nom, &prenom)
		if errScan == nil {
			fmt.Println(fmt.Sprintf("ID : %d, Nom : %s, Prénom: %s", id, nom, prenom))
		} else {
			err = errScan
		}
	}

	if err != nil {
		log.Fatal(err.Error())
	}

}
