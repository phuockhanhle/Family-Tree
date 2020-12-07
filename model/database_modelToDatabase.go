package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Insert_person(p *Person) error {
	if p.Birthday.IsZero() {
		_, err := insertPerson.Exec(p.FirstName, p.LastName, string(p.Gender), p.Rank, "0000-00-00")
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		_, err := insertPerson.Exec(p.FirstName, p.LastName, string(p.Gender), p.Rank, TimeToString(p.Birthday)[0:10])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	// dirty trick
	p.ID = GetNumberPerson()
	return nil
}

func Insert_1st_person(p *Person) error {
	err := Insert_person(p)

	err = InsertTree(1)

	err = SetFatherTree(1, 1)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Insert_nth_person(p_old *Person, p_new *Person) {
	_ = Insert_person(p_new)

	ID_new := GetIdByInfo(p_new.FirstName, p_new.LastName, p_new.Birthday)
	ID_old := GetIdByInfo(p_old.FirstName, p_old.LastName, p_old.Birthday)

	type_relation := GetRelation(p_new, p_old)
	_ = Insert_relation(ID_new, ID_old, type_relation)

	ID_father_tree, _ := GetIdFatherTree(ID_old)
	ID_mother_tree, _ := GetIdMotherTree(ID_old)

	switch GetRelation(p_new, p_old) {
	case ChildRole:
		if p_old.Gender == Male {
			_ = SetFatherTree(ID_new, ID_father_tree)
		} else {
			_ = SetMotherTree(ID_new, ID_mother_tree)
		}
	case ParentRole:
		if p_new.Gender == Male {
			_ = SetFatherTree(ID_new, ID_father_tree)
			_ = UpdateTreeRoot(ID_father_tree, ID_new)
		} else {
			_ = InsertTree(ID_new)
			id_tree := GetIdTreeByRoot(ID_new)
			_ = SetFatherTree(ID_new, id_tree)
			_ = SetMotherTree(ID_old, id_tree)
		}
	case SpouseRole:
		_ = InsertTree(ID_new)
		id_tree := GetIdTreeByRoot(ID_new)
		_ = SetFatherTree(ID_new, id_tree)
	}
}

// Macro for database relationship :
// If relation is spousal, husband is always on left columns (because wife is always right)
// If relation is parental, parent is always on first columns
func Insert_relation(id_source int, id_dest int, type_relation Role) error {
	//case relation is parental

	switch type_relation {

	case ChildRole:
		return Insert_relation(id_dest, id_source, ParentRole)
	case ParentRole:
		_, err := insertRelation.Exec(id_source, id_dest, "parental")
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	case SpouseRole:
		tmp, _ := GetPersonById(id_source)
		if tmp.Gender == Male {
			return Insert_relation(id_dest, id_source, SpouseRole)
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

func MakeRelationBetweenPeopleAlreadyInDB(id_source int, id_dest int, type_relation Role) error {
	_ = Insert_relation(id_source, id_dest, type_relation)

	ID_father_tree_source, _ := GetIdFatherTree(id_source)
	ID_father_tree_dest, _ := GetIdFatherTree(id_dest)
	switch type_relation {

	case ChildRole:
		tmp, _ := GetPersonById(id_dest)
		if tmp.Gender == Female {
			_ = SetMotherTree(id_source, ID_father_tree_dest)
		} else {
			_ = SetFatherTree(id_source, ID_father_tree_source)
		}
		return nil
	case ParentRole:
		_ = MakeRelationBetweenPeopleAlreadyInDB(id_dest, id_source, ChildRole)
		return nil
	case SpouseRole:
		//nothing to do because
		return nil
	case NilRole:
	}
	return nil

}

func InsertTree(id_root int) error {
	_, err := insertTree.Exec(id_root)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
