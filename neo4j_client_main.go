package main

import (
	"fmt"

	model "github.com/phuockhanhle/familytree/model"
)

func main() {
	uri := string("bolt://localhost:11006")
	username := string("neo4j")
	password := string("1234")

	vis := model.Person{
		ID:        model.GenerateID(),
		IDTree:    model.GenerateID(),
		FirstName: "Thanh Trung",
		LastName:  "Dinh",
		NickName:  "vis",
		Gender:    model.Male,
		HasChild:  false,
	}

	ba_phuong := model.Person{
		ID:        model.GenerateID(),
		IDTree:    vis.IDTree,
		FirstName: "Thanh Phuong",
		LastName:  "Dinh",
		NickName:  "nheo",
		Gender:    model.Male,
		HasChild:  true,
	}

	// vis.ID = "b34041b9-bda6-11eb-8c9d-b06ebfab25a7"

	var driver model.Neo4jDriver

	driver.CreateInstance(uri, username, password)
	defer driver.Close()

	vis = driver.RunTransaction(model.InsertPerson, vis).(model.Person)
	ba_phuong = driver.RunTransaction(model.InsertPerson, ba_phuong).(model.Person)

	driver.RunTransaction(model.InsertRelation, model.Relation{
		FromID:       vis.ID,
		ToID:         ba_phuong.ID,
		TypeRelation: "CHILDOF",
	})

	p := driver.RunTransaction(model.MatchPersonByID, vis.ID).(model.Person)
	fmt.Println(model.StructToMap(p))

	r := model.Relation{
		FromID:       vis.ID,
		TypeRelation: "CHILDOF",
	}
	list_p := driver.RunTransaction(model.MatchPersonByRelation, r).([]model.Person)
	fmt.Println(list_p)

	// if err == nil {
	// 	fmt.Println(greeting)
	// } else {
	// 	fmt.Println("failed!")
	// 	fmt.Println(err)
	// }
}
