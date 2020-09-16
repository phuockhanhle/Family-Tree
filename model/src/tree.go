package model

import (
	"encoding/json"
	"io/ioutil"
)

type Tree struct {
	ID       int
	Root     int
	Filename string
}

var TM []*Tree //Tree manager

func BuildTrees() {
	ClearTreeManager()
	PM.ReadFromCSV()
	for _, p := range PM.AllPeople {
		if p.IsRoot() {
			t := Tree{ID: len(TM) + 1, Root: p.ID, Filename: "FamilyTreeOf_" + p.LastName + ".json"}
			TM = append(TM, &t)
		}
	}
}

func (t Tree) WriteToJson() error {
	idRoot := t.Root
	root := PM.GetPeopleByID(idRoot)
	data := root.ToJSONForm()
	file, _ := json.MarshalIndent(data, "", "")
	_ = ioutil.WriteFile(t.Filename, file, 0644)
	return nil
}

func (t *Tree) UpdateRoot() error {
	return nil
}

func ClearTreeManager() {

}
