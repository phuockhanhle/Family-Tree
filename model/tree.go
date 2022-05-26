package model

import (
	json "encoding/json"
)

type Tree struct {
	Node Person 		`json:"node"`
	Children []Tree		`json:"children"`
}

func (tree Tree)ToJson() string {
	jsonTree, _ := json.Marshal(tree)
	return string(jsonTree)
}
