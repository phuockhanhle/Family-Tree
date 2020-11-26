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
var selectPersonById *sql.Stmt
var selectFatherTreePerson *sql.Stmt
var selectMotherTreePerson *sql.Stmt
var selectNumberPerson *sql.Stmt
var selectIdTreeByRoot *sql.Stmt
var selectRootByIdTree *sql.Stmt

// prepared statements for update or set
var updateFatherTree *sql.Stmt
var updateMotherTree *sql.Stmt
var updateTreeRootID *sql.Stmt
var updateBirthdayPerson *sql.Stmt
var updateDeathdayPerson *sql.Stmt

// prepared statements for insert
var insertPerson *sql.Stmt
var insertRelation *sql.Stmt
var insertTree *sql.Stmt

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

	selectPersonById, err = db.Prepare("select * from Person P WHERE P.ID_person = ? ")

	selectNumberPerson, err = db.Prepare("select count(*) from Person")

	updateFatherTree, err = db.Prepare("UPDATE Person SET ID_fatherTree = ? WHERE ID_person = ? ")

	updateMotherTree, err = db.Prepare("UPDATE Person SET ID_motherTree = ? WHERE ID_person = ? ")

	updateTreeRootID, err = db.Prepare("UPDATE Tree SET ID_root = ? WHERE ID_tree = ?")

	insertTree, err = db.Prepare("INSERT INTO TREE(ID_root) values ?")

	selectIdTreeByRoot, err = db.Prepare("SELECT ID_tree from Tree WHERE ID_root = ?")

	selectRootByIdTree, err = db.Prepare("SELECT ID_root from Tree WHERE ID_tree = ?")

	selectFatherTreePerson, err = db.Prepare("SELECT ID_FatherTree FROM Person WHERE ID_person = ?")

	selectMotherTreePerson, err = db.Prepare("SELECT ID_MotherTree FROM Person WHERE ID_person = ?")

	if err != nil {
		log.Fatal(err)
	}
}
