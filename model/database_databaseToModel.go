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
