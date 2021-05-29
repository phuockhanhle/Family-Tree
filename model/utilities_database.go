package model

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// // the database connection
// var db *sql.DB
// var err error

// // prepared statements for select
// var selectIDByName *sql.Stmt
// var selectAllPeople *sql.Stmt
// var selectParents *sql.Stmt
// var selectSpouse *sql.Stmt

// // prepared statements for insert
// var insertPerson *sql.Stmt
// var insertRelation *sql.Stmt

// func Connect_database() {
// 	db, err := sql.Open("mysql", "root:MeoMeo123!@#@(localhost:3306)/family_tree")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	err = db.Ping() // ensure alive...
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	insertPerson, err = db.Prepare("INSERT INTO Person(FirstName,LastName,Gender,Rank) VALUES (?, ?, ?, ?)")

// 	//create prepared statement for inserting a relation
// 	insertRelation, err = db.Prepare("INSERT INTO Relation(ID_source,ID_dest,type) VALUES (?, ?, ?)")

// 	//create prepared statement for getting ID of a Person by name
// 	selectIDByName, err = db.Prepare("SELECT ID_person FROM Person WHERE FirstName = ? And LastName = ? ")

// 	//create prepared statement for getting all people from database
// 	selectAllPeople, err = db.Prepare("SELECT * from Person")

// 	//create prepared statement for getting relation from database
// 	selectParents, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Parental' and R.ID_dest = ? ")

// 	//create prepared statement for getting relation from database
// 	selectSpouse, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Spousal' and (R.ID_dest = ? or R.ID_source = ? ) ")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func PrepareStatement() error {
// 	//create prepared statement for inserting a person
// 	insertPerson, err = db.Prepare("INSERT INTO Person(FirstName,LastName,Gender,Rank) VALUES (?, ?, ?, ?)")

// 	//create prepared statement for inserting a relation
// 	insertRelation, err = db.Prepare("INSERT INTO Relation(ID_source,ID_dest,type) VALUES (?, ?, ?)")

// 	//create prepared statement for getting ID of a Person by name
// 	selectIDByName, err = db.Prepare("SELECT ID_person FROM Person WHERE FirstName = ? And LastName = ? ")

// 	//create prepared statement for getting all people from database
// 	selectAllPeople, err = db.Prepare("SELECT * from Person")

// 	//create prepared statement for getting relation from database
// 	selectParents, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Parental' and R.ID_dest = ? ")

// 	//create prepared statement for getting relation from database
// 	selectSpouse, err = db.Prepare("SELECT R.ID_source FROM Person P, Relation R WHERE type = 'Spousal' AND R.ID_dest = ? ")

// 	return err
// }

// //------------------------ from model to database --------------------
// // func Insert_person(p *Person) error {
// // 	_, err = insertPerson.Exec(p.FirstName, p.LastName, string(p.Gender), p.Rank)

// // 	if err != nil {
// // 		log.Println(err)
// // 		return err
// // 	}
// // 	return nil
// // }

// // Macro for database relationship :
// // If relation is spousal, husband is always on left columns (because wife is always right)
// // If relation is parental, parent is always on first columns
// func Insert_relation(p_source *Person, p_dest *Person) error {
// 	//case relation is parental

// 	id_source, _ := GetIdByName(p_source.FirstName, p_source.LastName)
// 	id_dest, _ := GetIdByName(p_dest.FirstName, p_dest.LastName)

// 	switch GetRelation(p_source, p_dest) {

// 	case ChildRole:
// 		return Insert_relation(p_dest, p_source)
// 	case ParentRole:
// 		_, err = insertRelation.Exec(id_source, id_dest, "parental")
// 		if err != nil {
// 			log.Println(err)
// 			return err
// 		}
// 		return nil
// 	case SpouseRole:
// 		if p_source.Gender == Female {
// 			return Insert_relation(p_dest, p_source)
// 		} else {
// 			_, err = insertRelation.Exec(id_source, id_dest, "spousal")
// 			if err != nil {
// 				log.Println(err)
// 				return err
// 			}
// 			return nil
// 		}
// 	case NilRole:
// 		log.Println("not parental or spousal relations")
// 		return nil
// 	}

// 	return nil
// }

// func GetIdByName(FirstName string, LastName string) (int, error) {
// 	id_row, err := selectIDByName.Query(FirstName, LastName)
// 	if err != nil {
// 		log.Println("query error", err)
// 		return 0, err
// 	}

// 	defer id_row.Close()

// 	if !id_row.Next() {
// 		return 0, nil
// 	}

// 	var id int
// 	err = id_row.Scan(&id)
// 	if err != nil {
// 		log.Println("scan error", err)
// 		return 0, err
// 	}
// 	return id, nil

// }

// func Clear_tables() {
// 	var clearTablePerson *sql.Stmt
// 	var clearTableRelation *sql.Stmt

// 	clearTablePerson, err = db.Prepare("DELETE FROM Person")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	clearTableRelation, err = db.Prepare("DELETE FROM Relation")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	clearTableRelation.Exec()
// 	clearTablePerson.Exec()
// }

// //------------------------ from database to model --------------------
// func GetAllPeople() ([]*Person, error) {
// 	rows, err := selectAllPeople.Query()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	if !rows.Next() {
// 		return nil, nil
// 	}

// 	var res []*Person
// 	for {
// 		var tmp Person
// 		err = rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName,
// 			&tmp.NickName, &tmp.Gender, &tmp.Rank,
// 			&tmp.Birthday, &tmp.Deathday)
// 		res = append(res, &tmp)
// 		if !rows.Next() {
// 			break
// 		}
// 	}
// 	return res, nil

// }
