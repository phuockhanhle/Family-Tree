package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//------------------------ from model to database --------------------

func GetIdByName(FirstName string, LastName string) (int, error) {
	id_row, err := selectIDByName.Query(FirstName, LastName)
	if err != nil {
		log.Println("query error", err)
		return 0, err
	}

	defer id_row.Close()

	if !id_row.Next() {
		return 0, nil
	}

	var id int
	err = id_row.Scan(&id)
	if err != nil {
		log.Println("scan error", err)
		return 0, err
	}
	return id, nil
}

func Clear_tables() {
	var clearTablePerson *sql.Stmt
	var clearTableRelation *sql.Stmt

	clearTablePerson, err := db.Prepare("DELETE FROM Person")
	if err != nil {
		log.Fatal(err)
	}

	clearTableRelation, err = db.Prepare("DELETE FROM Relation")
	if err != nil {
		log.Fatal(err)
	}

	clearTableRelation.Exec()
	clearTablePerson.Exec()
}
