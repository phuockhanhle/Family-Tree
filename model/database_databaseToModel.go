package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPeople() ([]*Person, error) {
	rows, err := selectAllPeople.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	/*
		if !rows.Next() {
			return nil, nil
		}
	*/
	var res []*Person
	for rows.Next() {
		var tmp *Person

		tmp = assignPerson(rows)

		res = append(res, tmp)
	}
	return res, nil
}

func GetParentsFromDatabase(p *Person) ([]int, error) {
	rows, err := selectParentsPerson.Query(p.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var res []int
	for {
		var tmp int
		err = rows.Scan(&tmp)
		res = append(res, tmp)
		if !rows.Next() {
			break
		}
	}
	return res, nil
}

func GetSpousesFromDatabase(p *Person) ([]int, error) {
	if p.Gender == Male {
		rows, _ := selectWifePerson.Query(p.ID)

		defer rows.Close()

		if !rows.Next() {
			return nil, nil
		}

		var res []int
		for {
			var tmp string
			_ = rows.Scan(&tmp)
			res = append(res, StringToInt(tmp))
			if !rows.Next() {
				break
			}
		}
		return res, nil
	} else {
		rows, _ := selectHusbandPerson.Query(p.ID)

		defer rows.Close()

		if !rows.Next() {
			return nil, nil
		}

		var res []int
		for {
			var tmp int
			_ = rows.Scan(&tmp)
			res = append(res, tmp)
			if !rows.Next() {
				break
			}
		}
		return res, nil
	}
}

func GetChildrenFromDatabase(p *Person) ([]int, error) {
	rows, err := selectChildrenPerson.Query(p.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var res []int
	for {
		var tmp int
		err = rows.Scan(&tmp)
		res = append(res, tmp)
		if !rows.Next() {
			break
		}
	}
	return res, nil
}

func UpdateRelationPerson(p *Person) {
	p.Children, _ = GetChildrenFromDatabase(p)
	p.Spouse, _ = GetSpousesFromDatabase(p)
}

func UpdateRelation(people []*Person) {
	for _, p := range people {
		UpdateRelationPerson(p)
	}
}

func GetPersonById(ID_person int) (*Person, error) {
	rows, err := selectPersonById.Query(ID_person)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	res := assignPerson(rows)
	return res, nil
}

func assignPerson(rows *sql.Rows) *Person {
	var p Person

	var ID string
	var FirstName string
	var LastName string
	var NickName sql.NullString
	var Rank string
	var Birthday string
	var Deathday sql.NullString
	var Gender string

	_ = rows.Scan(&ID, &FirstName, &LastName,
		&NickName, &Gender, &Rank,
		&Birthday, &Deathday)

	p.ID = StringToInt(ID)
	p.FirstName = FirstName
	p.LastName = LastName
	p.NickName = NickName.String
	p.Rank = StringToInt(Rank)

	if Gender == "M" {
		p.Gender = Male
	} else {
		p.Gender = Female
	}
	p.Birthday = StringToTime(Birthday)
	p.Deathday = StringToTime(Deathday.String)

	return &p
}
