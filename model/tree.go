package model

// import (
// 	"encoding/csv"
// 	"errors"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// )

// /*
// //--------------------------------------- Old Design of Family Tree -------------------------------------------------

// type Tree struct {
// 	ID       int
// 	Root     int
// 	Filename string
// }

// var TM []*Tree //Tree manager

// func BuildTrees() {
// 	ClearTreeManager()
// 	PM.ReadFromCSV()
// 	for _, p := range PM.AllPeople {
// 		if p.IsRoot() {
// 			t := Tree{ID: len(TM) + 1, Root: p.ID, Filename: "FamilyTreeOf_" + p.LastName + ".json"}
// 			TM = append(TM, &t)
// 		}
// 	}
// }

// func (t Tree) WriteToJson() error {
// 	idRoot := t.Root
// 	root := PM.GetPeopleByID(idRoot)
// 	data := root.ToJSONForm()
// 	file, _ := json.MarshalIndent(data, "", "")
// 	_ = ioutil.WriteFile(t.Filename, file, 0644)
// 	return nil
// }

// func (t *Tree) UpdateRoot() error {
// 	return nil
// }

// func ClearTreeManager() {

// }
// */
// //--------------------------------------- New Design of Family Tree -------------------------------------------------

// type TreesManager struct {
// 	AllTrees []int
// }

// func (TM *TreesManager) GetRootOfTree(id int) int {
// 	return TM.AllTrees[id]
// }

// func (TM *TreesManager) AddTree(root int) int {
// 	TM.AllTrees = append(TM.AllTrees, root)
// 	return len(TM.AllTrees) - 1
// }

// func (TM *TreesManager) WriteToCSV(savePath string) error {
// 	csvFilePath := filepath.Join(savePath, "trees.csv")
// 	csvFile, err := os.Create(csvFilePath)
// 	defer csvFile.Close()
// 	if err != nil {
// 		return errors.New("TreeManager: error in open trees.csv")
// 	}
// 	writer := csv.NewWriter(csvFile)
// 	for idTree, treeRoot := range TM.AllTrees {
// 		writer.Write([]string{strconv.Itoa(idTree), strconv.Itoa(treeRoot)})
// 		writer.Flush()
// 	}
// 	return nil
// }

// var TM TreesManager

// type TreeGroups struct {
// 	FatherTree    int
// 	MotherTree    int
// 	InheritFather []int
// 	InheritMother []int
// }
// // -> fatherTrees(FatherTree + InheritFather)
// // -> motherTrees(MotherTree + InheritMother)

// // id, root, leaf

// func (g *TreeGroups) GetFatherTrees() []int {
// 	return append(g.InheritFather, g.FatherTree)
// }

// func (g *TreeGroups) GetMotherTrees() []int {
// 	return append(g.InheritMother, g.MotherTree)
// }

// func (g *TreeGroups) GetRoots() []int {
// 	var roots []int
// 	for _, tree := range append(g.GetFatherTrees(), g.GetMotherTrees()...) {
// 		roots = append(roots, TM.GetRootOfTree(tree))
// 	}
// 	return roots
// }

// func buildTreeGroups(TM *TreesManager) {

// 	for _, tree := range TM.AllTrees {
// 		// For each roots, goes down to each child???
// 	}

// }
