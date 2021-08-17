package model

import (
	"fmt"

	neo4j "github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jDriver struct {
	inst neo4j.Driver
}

func (neo4jDriver *Neo4jDriver) Close() {
	neo4jDriver.inst.Close()
}

func (neo4jDriver *Neo4jDriver) CreateInstance(uri, username, password string) {
	var err error
	neo4jDriver.inst, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
}

type TransactionOperation struct {
	access_mode      neo4j.AccessMode
	transaction_work neo4j.TransactionWork
}

func (neo4jDriver *Neo4jDriver) DeleteDatabase() {
	session := neo4jDriver.inst.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run("MATCH (n) DETACH DELETE n", nil)
		return nil, err
	})

	if err != nil {
		panic(err)
	}
}

func (neo4jDriver *Neo4jDriver) RunTransaction(trans_op func(param interface{}) TransactionOperation, param interface{}) interface{} {

	op := trans_op(param)

	session := neo4jDriver.inst.NewSession(neo4j.SessionConfig{AccessMode: op.access_mode})
	defer session.Close()

	transaction_func := session.ReadTransaction
	if op.access_mode == neo4j.AccessModeWrite {
		transaction_func = session.WriteTransaction
	}

	ret, err := transaction_func(op.transaction_work)

	if err != nil {
		panic(err)
	}

	return ret
}

func InsertPerson(p interface{}) TransactionOperation {
	person := p.(Person)

	return TransactionOperation{
		access_mode: neo4j.AccessModeWrite,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			ret, err := tx.Run("CREATE (p:Person $params) RETURN p", map[string]interface{}{"params": StructToMap(person)})

			if ret.Next() {
				props := ret.Record().Values[0].(neo4j.Node).Props
				p = MapToStruct(props, Person{}).(Person)
			}

			return p, err
		},
	}
}

func InsertRelation(r interface{}) TransactionOperation {
	relation := r.(Relation)

	params := map[string]interface{}{
		"FromID": relation.FromID,
		"ToID":   relation.ToID,
	}

	from_person := string("from_person")
	to_person := string("to_person")

	match_from_person := fmt.Sprintf("MATCH (%s:Person {ID: $FromID})", from_person)
	match_to_person := fmt.Sprintf("MATCH (%s:Person {ID: $ToID})", to_person)
	create_relation := fmt.Sprintf("CREATE (%s)-[:%s]->(%s)", from_person, relation.TypeRelation, to_person)

	return TransactionOperation{
		access_mode: neo4j.AccessModeWrite,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			_, err := tx.Run(fmt.Sprintf("%s %s %s", match_from_person, match_to_person, create_relation), params)
			return relation, err
		},
	}
}

func UpdatePerson(i interface{}) TransactionOperation {
	info := i.(InfoUpdate)

	params := map[string]interface{}{
		"PersonID": info.PersonID,
		"Field":    info.Field,
		"Value":    info.Value,
	}

	person := string("person")

	match_person := fmt.Sprintf("MATCH (%s:Person {ID: $PersonID})", person)
	update_info := fmt.Sprintf("SET %s.%s = $Value", person, info.Field)

	return TransactionOperation{
		access_mode: neo4j.AccessModeWrite,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			_, err := tx.Run(fmt.Sprintf("%s %s", match_person, update_info), params)
			return nil, err
		},
	}
}

func MatchPersonByID(person_id interface{}) TransactionOperation {

	return TransactionOperation{
		access_mode: neo4j.AccessModeRead,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			ret, err := tx.Run("MATCH (p:Person {ID: $ID}) RETURN p", map[string]interface{}{"ID": person_id.(string)})

			var p Person

			if ret.Next() {
				props := ret.Record().Values[0].(neo4j.Node).Props
				p = MapToStruct(props, Person{}).(Person)
			}

			if ret.Next() {
				panic("Multiple person detected for 1 id")
			}

			return p, err
		},
	}
}

func MatchPeopleByRelation(r interface{}) TransactionOperation {

	rel := r.(Relation)

	if rel.FromID != "" && rel.ToID != "" {
		panic("both nodes already exist")
	}

	from_person := "from_p: Person"
	if rel.FromID != "" {
		from_person += " {ID: $FromID}"
	}

	to_person := "to_p: Person"
	if rel.ToID != "" {
		to_person += " {ID: $ToID}"
	}

	var return_result string
	if rel.FromID != "" {
		return_result = "to_p"
	} else if rel.ToID != "" {
		return_result = "from_p"
	}

	query := fmt.Sprintf("MATCH (%s)-[:%s]->(%s) RETURN %s", from_person, rel.TypeRelation, to_person, return_result)

	return TransactionOperation{
		access_mode: neo4j.AccessModeRead,
		transaction_work: func(tx neo4j.Transaction) (interface{}, error) {
			ret, err := tx.Run(query, StructToMap(rel))

			var list_p []Person

			for ret.Next() {
				props := ret.Record().Values[0].(neo4j.Node).Props
				list_p = append(list_p, MapToStruct(props, Person{}).(Person))
			}

			return list_p, err
		},
	}
}
