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

type strDatabaseSetting struct {
	SettingName  string `json:"name"`
	DatabaseName string `json:"dbName"`
	HostName     string `json:"server"`
	Port         int    `json:"port"`
	Driver       string `json:"driver"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

type strContact struct {
	ID     int
	Nom    string
	Prenom string
}

func main() {

	var db *sql.DB
	var err error

	db, err = openDB()
	defer db.Close()
	if err == nil {
		var contact strContact

		contact.Nom = "Dupond"
		contact.Prenom = "Alain"

		err = createContact(db, &contact)
		if err == nil {
			err = listecontacts(db)
			if err == nil {
				err = getContact(db, 1)
			}
		}
	}

	if err != nil {
		log.Fatal(err.Error())
	}

}

func openDB() (*sql.DB, error) {
	var dbSetting strDatabaseSetting

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

func listecontacts(db *sql.DB) error {

	var contact strContact
	var lstContacts []strContact
	var err error

	rows, errQuery := db.Query("Select * from contact")
	if errQuery == nil {
		defer rows.Close()
		for rows.Next() {

			errScan := rows.Scan(&contact.ID, &contact.Nom, &contact.Prenom)
			if errScan == nil {
				// alimentation d'un taleau de contact
				lstContacts = append(lstContacts, contact)

				fmt.Println(fmt.Sprintf("ID : %d, Nom : %s, Prénom: %s", contact.ID, contact.Nom, contact.Prenom))
			} else {
				err = errScan
			}
		}
	} else {
		err = errQuery
	}
	return err
}

func getContact(db *sql.DB, id int) error {
	var contact strContact
	var err error

	row := db.QueryRow("SELECT nom,prenom FROM contact WHERE id = $1;", id)
	errScan := row.Scan(&contact.Nom, &contact.Prenom)
	if errScan == nil {
		fmt.Println(fmt.Sprintf("ID : %d, Nom : %s, Prénom: %s", id, contact.Nom, contact.Prenom))
	} else {
		err = errScan
	}

	return err
}

func createContact(db *sql.DB, contact *strContact) error {
	var err error

	dbtransact, errTransac := db.Begin()
	if err == nil {

		result, errExec := dbtransact.Exec("INSERT INTO contact (nom,prenom) values ($1,$2)", contact.Nom, contact.Prenom)
		if errExec == nil {
			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())
		} else {
			err = errExec
			dbtransact.Rollback()
		}

		dbtransact.Commit()
	} else {
		err = errTransac
	}

	return err
}

func updateContact(db *sql.DB, id int, contact *strContact) error {
	var err error
	dbtransact, errTransac := db.Begin()
	if err == nil {

		result, errExec := dbtransact.Exec("INSERT contact set nom=$1, prenom=$2 where id=$3 ", contact.Nom, contact.Prenom, id)
		if errExec == nil {
			fmt.Println(result.RowsAffected())
		} else {
			err = errExec
			dbtransact.Rollback()
		}

		dbtransact.Commit()
	} else {
		err = errTransac
	}

	return err
}

func deleteContact(db *sql.DB, id int) error {
	var err error
	dbtransact, errTransac := db.Begin()
	if err == nil {

		result, errExec := dbtransact.Exec("DELETE FROM contact where id=$1 ", id)
		if errExec == nil {
			fmt.Println(result.RowsAffected())
		} else {
			err = errExec
			dbtransact.Rollback()
		}

		dbtransact.Commit()
	} else {
		err = errTransac
	}

	return err
}
