package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//PeopleManager manage a group of people
type PeopleManager struct {
	AllPeople []*Person
}

// GetNrPeople returns the number of people in PeopleManager
func (pm PeopleManager) GetNrPeople() int {
	return len(pm.AllPeople)
}

// GetPeopleByID returns pointer to a person with his ID
func (pm PeopleManager) GetPeopleByID(ID int) *Person {
	if pm.GetNrPeople() > 0 {
		return pm.AllPeople[ID-1]
	}
	return nil
}

// AddNewPerson creates new person and adds to PeopleManager
func (pm *PeopleManager) AddNewPerson(firstName string, lastName string, gender GenderType) *Person {
	newPerson := Person{ID: pm.GetNrPeople() + 1, FirstName: firstName, LastName: lastName, Gender: gender,
		Birthday: time.Time{}, Deathday: time.Time{}, Rank: DefaultRank}
	if newPerson.ID == 1 {
		newPerson.Rank = 0
	}
	pm.AllPeople = append(pm.AllPeople, &newPerson)
	return &newPerson
}

// AddNewRelation adds a relation from fromID person to toID person
func (pm PeopleManager) AddNewRelation(fromID int, toID int, relaRole Role) {
	fromPerson := pm.GetPeopleByID(fromID)
	toPerson := pm.GetPeopleByID(toID)
	switch relaRole {
	case ParentRole:
		fromPerson.AddChildren(toPerson)
		toPerson.AddParent(fromPerson)
	case SpouseRole:
		fromPerson.AddSpouse(toPerson)
		toPerson.AddSpouse(fromPerson)
	case ChildRole:
		fromPerson.AddParent(toPerson)
		toPerson.AddChildren(fromPerson)
	}
	toPerson.UpdateRank(fromPerson, relaRole)
}

// PersonToRowEntry convert all information about one person into row entry
func PersonToRowEntry(p Person) []string {
	var spouseIDs []string
	for _, s := range p.Spouse {
		spouseIDs = append(spouseIDs, strconv.Itoa(s.ID))
	}
	var dataRow []string
	dataRow = append(dataRow, strconv.Itoa(p.ID))
	dataRow = append(dataRow, p.FirstName)
	dataRow = append(dataRow, p.LastName)
	dataRow = append(dataRow, string(p.Gender))
	dataRow = append(dataRow, TimeToString(p.Birthday))
	dataRow = append(dataRow, TimeToString(p.Deathday))
	dataRow = append(dataRow, strconv.Itoa(p.Rank))
	dataRow = append(dataRow, strconv.Itoa(p.GetDadID()))
	dataRow = append(dataRow, strconv.Itoa(p.GetMomID()))
	dataRow = append(dataRow, strings.Join(spouseIDs, IDSeparator))
	return dataRow
}

// WriteToCSV write all entries in PeopleManager into a csv file named people.csv
func (pm PeopleManager) WriteToCSV(savePath string) error {
	csvFile := filepath.Join(savePath, "people.csv")
	f, err := os.Create(csvFile)
	defer f.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		return errors.New("error in open file")
	}
	w := csv.NewWriter(f)
	for _, p := range pm.AllPeople {
		dataRow := PersonToRowEntry(*p)
		w.Write(dataRow)
		w.Flush()
	}
	return nil
}

// Clear method deletes all data in PeopleManager
func (pm *PeopleManager) Clear() {
	// to test if gabage collector runs
	//pm.AllPeople = append(pm.AllPeople[:0], pm.AllPeople[pm.GetNbr():]...)
	pm.AllPeople = nil // I think that we should do like this to allow gabage collector to delete memories
}

// ReadFromCSV reads csv file and create new people based on those data
func (pm *PeopleManager) ReadFromCSV() {
	// Open csv file
	fmt.Println("starting read")
	f, err := os.Open("people.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Clear currentData in PeopleManager
	pm.Clear()
	// List to store the relations of all people
	var listPersoRela []PersoRela
	// Read and parse csv file
	reader := csv.NewReader(f)
	for {
		// Read csv until end-of-file
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Create new person
		p := pm.AddNewPerson(record[1], record[2], GenderType(record[3][0]))
		// Update birthday, deathday and rank
		p.Birthday = StringToTime(record[4])
		p.Deathday = StringToTime(record[5])
		p.Rank = StringToInt(record[6])
		// Create dict of relations (dad, mom, spouse, children)
		var pRela = make(PersoRela)
		var listParentIDs []int
		for _, ID := range []int{StringToInt(record[7]), StringToInt(record[8])} {
			if ID != IDNotFound {
				listParentIDs = append(listParentIDs, ID)
			}
		}
		if listParentIDs != nil {
			pRela[ParentRole] = listParentIDs
		}
		var listSpouseIDs []int
		for _, stringID := range strings.Split(record[9], IDSeparator) {
			if StringToInt(stringID) != IDNotFound {
				listSpouseIDs = append(listSpouseIDs, StringToInt(stringID))
			}
		}
		if listSpouseIDs != nil {
			pRela[SpouseRole] = listSpouseIDs
		}
		// Append dict of relations of each person to listPersoRela
		listPersoRela = append(listPersoRela, pRela)
	}
	// Update personal relations
	for i, pRela := range listPersoRela {
		for relaRole, toIDs := range pRela {
			for _, toID := range toIDs {
				// AddParent also adds children, so no need to do that alone
				pm.AddNewRelation(pm.AllPeople[i].ID, toID, relaRole)
			}
		}
	}
}
