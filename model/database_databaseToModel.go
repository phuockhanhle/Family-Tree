package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPeople() ([]*Person, error) {
	rows, err := selectAllPeople.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var res []*Person
	for {
		var tmp Person

		err = rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName,
			&tmp.NickName, &tmp.Gender, &tmp.Rank,
			&tmp.Birthday, &tmp.Deathday)

		res = append(res, &tmp)

		if !rows.Next() {
			break
		}
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
	rows, err := selectSpousesPerson.Query(p.ID, p.ID)
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

func UpdateRelationFromDB(p *Person) {
	p.Children, _ = GetChildrenFromDatabase(p)
	p.Spouse, _ = GetSpousesFromDatabase(p)
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
	var res Person
	var ID string
	var FirstName string
	var LastName string
	var NickName string
	var Rank string
	var Birthday string
	var Deathday string
	var Gender string

	err = rows.Scan(&ID, &FirstName, &LastName,
		&NickName, Gender, &Rank,
		&Birthday, &Deathday)

	res.ID = StringToInt(ID)
	res.FirstName = FirstName
	res.LastName = LastName
	res.NickName = NickName
	res.Rank = StringToInt(Rank)

	if Gender == "M" {
		res.Gender = Male
	} else {
		res.Gender = Female
	}
	res.Birthday = StringToTime(Birthday)
	res.Deathday = StringToTime(Deathday)
	return &res, nil
}
