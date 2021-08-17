package model

type Controller struct {
	Driver Neo4jDriver
}

func (controller *Controller) InitDatabase() {
	uri := string("bolt://localhost:11006")
	username := string("neo4j")
	password := string("1234")

	controller.Driver.CreateInstance(uri, username, password)
	// defer controller.Driver.Close()
}

func insertChildOfRelation(Driver *Neo4jDriver, childID string, parentID string) interface{} {
	ret := Driver.RunTransaction(InsertRelation, Relation{
		FromID:       childID,
		ToID:         parentID,
		TypeRelation: "CHILDOF",
	})

	return ret
}

func (controller *Controller) InsertChildFromDad(child Person, dad Person) {
	child.IDTree = dad.IDTree

	_ = controller.Driver.RunTransaction(InsertPerson, child).(Person)
	_ = insertChildOfRelation(&controller.Driver, child.ID, dad.ID)

	_ = controller.Driver.RunTransaction(UpdatePerson, InfoUpdate{
		PersonID: dad.ID,
		Field:    "HasChild",
		Value:    true,
	})
}

func (controller *Controller) InsertChildFromMom(child Person, mom Person) {
	child.IDTree = GenerateID()

	_ = controller.Driver.RunTransaction(InsertPerson, child)
	_ = insertChildOfRelation(&controller.Driver, child.ID, mom.ID)
	_ = controller.Driver.RunTransaction(UpdatePerson, InfoUpdate{
		PersonID: mom.ID,
		Field:    "HasChild",
		Value:    true,
	})
}

func (controller *Controller) InsertDadFromChild(dad Person, child Person) {
	dad.IDTree = child.IDTree
	dad.HasChild = true

	_ = controller.Driver.RunTransaction(InsertPerson, dad).(Person)
	_ = insertChildOfRelation(&controller.Driver, child.ID, dad.ID)
}

func (controller *Controller) InsertMomFromChild(mom Person, child Person) {
	mom.IDTree = GenerateID()
	mom.HasChild = true

	_ = controller.Driver.RunTransaction(InsertPerson, mom)
	_ = insertChildOfRelation(&controller.Driver, child.ID, mom.ID)
}

// func (controller *Controller) DeleteDatabase() {
// 	controller.Driver.RunTransaction(DeleteDatabase, nil)
// }
