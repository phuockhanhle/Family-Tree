package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// the database connection
var db *sql.DB

// prepared statements for select
var selectIDByName *sql.Stmt
var selectAllPeople *sql.Stmt
var selectParentsPerson *sql.Stmt
var selectSpousesPerson *sql.Stmt
var selectChildrenPerson *sql.Stmt
var selectAllRoot *sql.Stmt

// prepared statements for insert
var insertPerson *sql.Stmt
var insertRelation *sql.Stmt

func Connect_database() {
	db, err := sql.Open("mysql", "root:MeoMeo123!@#@(localhost:3306)/family_tree")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping() // ensure alive...
	if err != nil {
		log.Fatal(err)
	}

	insertPerson, err = db.Prepare("INSERT INTO Person(FirstName,LastName,Gender,Rank) VALUES (?, ?, ?, ?)")

	//create prepared statement for inserting a relation
	insertRelation, err = db.Prepare("INSERT INTO Relation(ID_source,ID_dest,type) VALUES (?, ?, ?)")

	//create prepared statement for getting ID of a Person by name
	selectIDByName, err = db.Prepare("SELECT ID_person FROM Person WHERE FirstName = ? And LastName = ? ")

	//create prepared statement for getting all people from database
	selectAllPeople, err = db.Prepare("SELECT * from Person")

	//create prepared statement for getting relation from database
	selectParentsPerson, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Parental' and R.ID_dest = ? ")

	//create prepared statement for getting relation from database
	selectSpousesPerson, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Spousal' and (R.ID_dest = ? or R.ID_source = ? ) ")

	//create prepared statement for getting relation from database
	selectChildrenPerson, err = db.Prepare("SELECT R.ID_dest FROM Person P, Relation R WHERE type = 'Parental' and R.ID_source = ? ")

	selectAllRoot, err = db.Prepare("select P.ID_person from Person P LEFT JOIN Relation R on R.ID_dest = P.ID_person WHERE R.ID_dest IS NULL OR type = 'spousal'")

	if err != nil {
		log.Fatal(err)
	}
}
