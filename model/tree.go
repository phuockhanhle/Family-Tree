package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Tree struct {
	Id            int
	Root          int
	Filename_json string
}

type People_Json_Form struct {
	ID_self     int
	ID_children []*People_Json_Form
}

var TM []*Tree //Tree manager

func Build_trees(Path_CSV_file string) {
	Clear_TreeManager()
	PM.Read_CSV()
	for i := 0; i < PM.GetNbr(); i++ {
		if PM.AllPeople[i].Dad == nil && PM.AllPeople[i].Mom == nil {
			tmp := &Tree{Root: PM.AllPeople[i].Id, Filename_json: "Tree_root_" + PM.AllPeople[i].Ten + ".json"}
			AddTree(tmp)
			tmp.Id = len(TM)
		}
	}
}
func (p *People) To_Json_Form() *People_Json_Form {

	res := &People_Json_Form{ID_self: p.Id}
	children := p.GetChildren()
	if len(children) == 0 {
		return res
	}
	if len(children) > 0 {
		for i := 0; i < len(children); i++ {
			tmp := children[i].To_Json_Form()
			res.ID_children = append(res.ID_children, tmp)
		}
	}
	return res
}

func (t Tree) Savefile_Json() error {
	id_root := t.Root
	root := PM.GetPeopleByIndex(id_root)
	data := root.To_Json_Form()
	file, _ := json.MarshalIndent(data, "", "")
	_ = ioutil.WriteFile(t.Filename_json, file, 0644)
	return nil
}

func (t *Tree) Update_Root() error {
	return nil
}

func AddTree(t *Tree) error {
	for i := 0; i < len(TM); i++ {
		if t.Id == TM[i].Id {
			return errors.New("this tree already existes")
		}
	}
	TM = append(TM, t)
	return nil
}

func Clear_TreeManager() {

}
