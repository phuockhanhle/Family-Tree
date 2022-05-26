package model

import (
	"fmt"
	// json "encoding/json"
	neo4j "github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetTreeByID(IDTree interface{}) TransactionOperation {

	params := map[string]interface{}{
		"IDTree": IDTree.(string),
	}

	query_match := "Match p = (root:Person{IDTree: $IDTree})-[:PARENTOF*]->(:Person{IDTree: $IDTree})"
	query_filter := "WHERE NOT exists(()-[:PARENTOF]->(root)) WITH COLLECT(p) as ps"
	query_convert_tree := "CALL apoc.convert.toTree(ps) yield value RETURN value"

	return TransactionOperation{
		access_mode: neo4j.AccessModeRead,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			ret, err := tx.Run(fmt.Sprintf("%s %s %s", query_match, query_filter, query_convert_tree), params)

			var trees []Tree

			if ret.Next() {
				// props := ret.Record().Values[0].(neo4j.Node).Props
				// tree = MapToStruct(ret.Record().Values[0].(map[string]interface{}), Tree{}).()
				tree := constructTree(ret.Record().Values[0].(map[string]interface{}))
				trees = append(trees, tree)
			}

			return trees[0], err
		},
	}
}

func constructTree(queryTreeResult map[string]interface{}) Tree {

	person := make(map[string]interface{})
	var children []Tree

	for key, value := range queryTreeResult {
		if key == "parentof" {
			for _, subTree := range value.([]interface{}) {
				children = append(children, constructTree(subTree.(map[string]interface{})))
			}
		} else if key != "_type" && key != "_id" {
			person[key] = value
		}
	}

	return Tree{Node: MapToStruct(person, Person{}).(Person), Children: children}
}