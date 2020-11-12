package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Insert_person(p *Person) error {
	_, err := insertPerson.Exec(p.FirstName, p.LastName, string(p.Gender), p.Rank)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Macro for database relationship :
// If relation is spousal, husband is always on left columns (because wife is always right)
// If relation is parental, parent is always on first columns
func Insert_relation(p_source *Person, p_dest *Person) error {
	//case relation is parental

	id_source, _ := GetIdByName(p_source.FirstName, p_source.LastName)
	id_dest, _ := GetIdByName(p_dest.FirstName, p_dest.LastName)

	switch GetRelation(p_source, p_dest) {

	case ChildRole:
		return Insert_relation(p_dest, p_source)
	case ParentRole:
		_, err := insertRelation.Exec(id_source, id_dest, "parental")
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	case SpouseRole:
		if p_source.Gender == Female {
			return Insert_relation(p_dest, p_source)
		} else {
			_, err := insertRelation.Exec(id_source, id_dest, "spousal")
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	case NilRole:
		log.Println("not parental or spousal relations")
		return nil
	}

	return nil
}
