package test

import (
	model "github.com/phuockhanhle/familytree/model"
)

type TestCreateNode struct {
	info string
	toInsert model.Person,
	fromPerson: model.Person,
	refOutput: model.Person
}

type TestMatchNode struct {
	info: string,
	input: MatchProperties,
	refOutput: Person,
}

// insert child -> update : 
// 	- child.id_tree (copy from dad or create new from mom) 
//  - parent.HasChild

var createNodeTestSuite = []TestCreateNode [
	TestCreateNode{
		info: "insert child from dad"
		toInsert: ...
		fromPerson: ...
		refOutput: [IDTree(child) == IDTree(dad) && HasChild(dad) == true]
	},
	TestCreateNode{
		info: "insert child from mom"
		toInsert: ...
		fromPerson: ...
		refOutput: [IDTree(child) != IDTree(mom) && HasChild(mom) == true]
	},
	TestCreateNode{
		info: "insert dad from child"
		toInsert: ...
		fromPerson: ...
		refOutput: [IDTree(dad) == IDTree(child) && HasChild(dad) == true]
	},
	TestCreateNode{
		info: "insert mom from child"
		toInsert: ...
		fromPerson: ...
		refOutput: [IDTree(mom) != IDTree(child) && HasChild(mom) == true]
	},
	TestCreateNode{
		info: "insert spouse (2 subcases: spouse with HasChild = true vs HasChild = false)"
		toInsert: ...
		fromPerson: ...
		refOutput: [HasChild(fromPerson) == hasChild(toInsert)]
	},
]

var matchNodeTestSuite = []TestMatchNode [
	TestMatchNode{
		info: "match people by idtree"
		input: ...
		refOutput: ...
	},
]

func test(testData: TestCreateNode) {
	for _, data := createNodeTestSuite {
		output = insert(toInsert, fromPerson)
		assert(output == refOutput, info)
	}
}

