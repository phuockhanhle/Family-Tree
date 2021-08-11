package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	model "github.com/phuockhanhle/familytree/model"
)

func main() {
	uri := string("bolt://localhost:11006")
	username := string("neo4j")
	password := string("1234")

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
	// 	IDTree:    vis.IDTree,
	// 	FirstName: "Thanh Phuong",
	// 	LastName:  "Dinh",
	// 	NickName:  "nheo",
	// 	Gender:    model.Male,
	// 	HasChild:  true,
	// }

	var driver model.Neo4jDriver

	driver.CreateInstance(uri, username, password)
	defer driver.Close()

	csvFile, err := os.Open("./data/people.csv")
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

		gender := model.Male

		if line[3] == "F" {
			gender = model.Female
		}

		hasChild, _ := strconv.ParseBool(line[4])

		tmp := model.Person{
			FirstName: line[0],
			LastName:  line[1],
			NickName:  line[2],
			Gender:    gender,
			HasChild:  hasChild,
		}

		tmp = driver.RunTransaction(model.InsertPerson, tmp).(model.Person)
	}
}
