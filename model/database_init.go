package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// the database connection
var db *sql.DB

// prepared statements for select
var selectIDByInfo *sql.Stmt
var selectIDByName *sql.Stmt
var selectAllPeople *sql.Stmt
var selectParentsPerson *sql.Stmt
var selectHusbandPerson *sql.Stmt
var selectWifePerson *sql.Stmt
var selectChildrenPerson *sql.Stmt
var selectAllRoot *sql.Stmt
var selectPersonById *sql.Stmt
var selectFatherTreePerson *sql.Stmt
var selectMotherTreePerson *sql.Stmt
var selectNumberPerson *sql.Stmt
var selectIdTreeByRoot *sql.Stmt
var selectRootByIdTree *sql.Stmt

var selectGenderById *sql.Stmt

// prepared statements for update or set
var updateFatherTree *sql.Stmt
var updateMotherTree *sql.Stmt
var updateTreeRootID *sql.Stmt
var updateBirthdayPerson *sql.Stmt
var updateDeathdayPerson *sql.Stmt

// need solution
var updateFatherTreeByValueOfFatherTree *sql.Stmt

// prepared statements for insert
var insertPerson *sql.Stmt
var insertRelation *sql.Stmt
var insertTree *sql.Stmt

// prepared statements for delete

var deleteTree *sql.Stmt

func Connect_database() {
	db, err := sql.Open("mysql", "root:MeoMeo123!@#@(localhost:3306)/family_tree?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping() // ensure alive...
	if err != nil {
		log.Fatal(err)
	}

	insertPerson, err = db.Prepare("INSERT INTO Person(FirstName,LastName,Gender,Rank,Birthday) VALUES (?, ?, ?, ?, ?)")

	insertTree, err = db.Prepare("INSERT INTO Tree(ID_root) VALUES (?) ")
	//create prepared statement for inserting a relation
	insertRelation, err = db.Prepare("INSERT INTO Relation(ID_source,ID_dest,type) VALUES (?, ?, ?)")

	//create prepared statement for getting ID of a Person by name
	selectIDByInfo, err = db.Prepare("SELECT ID_person FROM Person WHERE FirstName = ? And LastName = ? And Birthday = ?")
	selectIDByName, err = db.Prepare("SELECT ID_person FROM Person WHERE FirstName = ? And LastName = ?")

	//create prepared statement for getting all people from database
	selectAllPeople, err = db.Prepare("SELECT ID_person,FirstName,LastName,NickName, Gender, Rank, Birthday, Deathday from Person")

	//create prepared statement for getting relation from database
	selectParentsPerson, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Parental' and R.ID_dest = ? ")

	//create prepared statement for getting relation from database
	selectHusbandPerson, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Spousal' and R.ID_dest = ? GROUP BY R.ID_source")
	selectWifePerson, err = db.Prepare("SELECT R.ID_dest FROM Person P, Relation R WHERE type = 'Spousal' and R.ID_source = ? GROUP BY R.ID_dest")

	//create prepared statement for getting relation from database
	selectChildrenPerson, err = db.Prepare("SELECT R.ID_dest FROM Person P, Relation R WHERE type = 'Parental' and R.ID_source = ? GROUP BY R.ID_dest")

	selectAllRoot, err = db.Prepare("select P.ID_person from Person P LEFT JOIN Relation R on R.ID_dest = P.ID_person WHERE R.ID_dest IS NULL OR type = 'spousal' ")

	selectPersonById, err = db.Prepare("select ID_person,FirstName,LastName,NickName, Gender, Rank, Birthday, Deathday from Person P WHERE P.ID_person = ? ")

	selectGenderById, err = db.Prepare("select Gender from Person Where ID_person = ?")

	selectNumberPerson, err = db.Prepare("select count(*) from Person")

	updateFatherTree, err = db.Prepare("UPDATE Person SET ID_fatherTree = ? WHERE ID_person = ? ")

	updateMotherTree, err = db.Prepare("UPDATE Person SET ID_motherTree = ? WHERE ID_person = ? ")

	updateTreeRootID, err = db.Prepare("UPDATE Tree SET ID_root = ? WHERE ID_tree = ?")

	updateFatherTreeByValueOfFatherTree, err = db.Prepare("UPDATE Person SET ID_fatherTree = ? WHERE ID_FatherTree = ?")

	selectIdTreeByRoot, err = db.Prepare("SELECT ID_tree from Tree WHERE ID_root = ?")

	selectRootByIdTree, err = db.Prepare("SELECT ID_root from Tree WHERE ID_tree = ?")

	selectFatherTreePerson, err = db.Prepare("SELECT ID_FatherTree FROM Person WHERE ID_person = ?")

	selectMotherTreePerson, err = db.Prepare("SELECT ID_MotherTree FROM Person WHERE ID_person = ?")

	deleteTree, err = db.Prepare("DELETE FROM Tree WHERE ID_tree = ?")
	if err != nil {
		log.Fatal(err)
	}
}
