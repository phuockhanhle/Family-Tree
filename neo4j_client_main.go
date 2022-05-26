package main

import (
	"os"
	"github.com/phuockhanhle/familytree/model"
	"fmt"
	"encoding/csv"
	"strconv"
)

func addPersonByCsv(driver model.Neo4jDriver) map[string]model.Person {
	csvFile, err := os.Open("./data/people.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	var result = make(map[string]model.Person)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {

		gender := model.Male

		if line[3] == "F" {
			gender = model.Female
		}

		hasChild, _ := strconv.ParseBool(line[4])

		tmp := model.Person{
			ID:        model.GenerateID(),
			IDTree:    line[5],
			FirstName: line[0],
			LastName:  line[1],
			NickName:  line[2],
			Gender:    gender,
			HasChild:  hasChild,
		}

		person := driver.RunTransaction(model.InsertPerson, tmp).(model.Person)
		result[person.FirstName] = person
	}
	return result
}

func addRelationByCsv(driver model.Neo4jDriver, mapPersonByID map[string]model.Person) {
	csvFile, err := os.Open("./data/relation.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	

	for _, line := range csvLines {

		fromPerson := mapPersonByID[line[0]]
		toPerson := mapPersonByID[line[1]]

		_ = driver.RunTransaction(model.InsertRelation, model.Relation{
			FromID:       fromPerson.ID,
			ToID:         toPerson.ID,
			TypeRelation: line[2],
		})

		// if toPerson.Gender == model.Male {
		// 	_ = driver.RunTransaction(model.UpdatePerson, model.InfoUpdate{
		// 		PersonID: 	toPerson.ID,
		// 		Field: 		"IDTree",
		// 		Value: 		fromPerson.IDTree,
		// 	})
		// }		
	}
}


func main() {
	uri := string("bolt://localhost:7687")
	username := string("neo4j")
	password := string("1234")

	var driver model.Neo4jDriver

	driver.CreateInstance(uri, username, password)
	defer driver.Close()

	driver.DeleteDatabase()

	var mapPersonByID = addPersonByCsv(driver)
	addRelationByCsv(driver, mapPersonByID)

	// var controller model.Controller

	// controller.InitDatabase()

	// vis := model.Person{
	// 	ID:        model.GenerateID(),
	// 	IDTree:    model.GenerateID(),
	// 	FirstName: "Thanh Trung",
	// 	LastName:  "Dinh",
	// 	NickName:  "vis",
	// 	Gender:    model.Male,
	// 	HasChild:  false,
	// }

	// ba_phuong := model.Person{
	// 	ID:        model.GenerateID(),
	// 	IDTree:    model.GenerateID(),
	// 	FirstName: "Thanh Phuong",
	// 	LastName:  "Dinh",
	// 	NickName:  "",
	// 	Gender:    model.Male,
	// 	HasChild:  false,
	// }

	// controller.Driver.RunTransaction(model.InsertPerson, ba_phuong)

	// controller.InsertChildFromDad(vis, ba_phuong)
}
