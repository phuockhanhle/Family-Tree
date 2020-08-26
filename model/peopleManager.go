package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

//manage a group of people
type PeopleManager struct {
	AllPeople []*People
}

func (pm PeopleManager) GetNbr() int {
	return len(pm.AllPeople)
}

func (pm PeopleManager) GetPeopleByIndex(Id int) *People {
	if pm.GetNbr() > 0 {
		for i := 0; i < pm.GetNbr(); i++ {
			if pm.AllPeople[i].Id == Id {
				return pm.AllPeople[i]
			}
		}
	}
	return nil
}

func (pm *PeopleManager) AddPeople(p *People) error {
	if pm.GetPeopleByIndex(p.Id) != nil {
		return errors.New("this person already existes")
	}
	pm.AllPeople = append(pm.AllPeople, p)
	return nil
}

//make data to save in a csv file
func (pm PeopleManager) CreateData() [][]string {
	var res [][]string
	for i := 0; i < pm.GetNbr(); i++ {
		tmp := *(pm.AllPeople[i])
		var dataRow []string
		dataRow = append(dataRow, strconv.Itoa(tmp.Id))
		dataRow = append(dataRow, tmp.Ten)
		dataRow = append(dataRow, tmp.Ho)
		if tmp.DayBirth.IsZero() {
			dataRow = append(dataRow, "nil")
		} else {
			dataRow = append(dataRow, tmp.DayBirth.String())
		}
		if tmp.DayDeath.IsZero() {
			dataRow = append(dataRow, "nil")
		} else {
			dataRow = append(dataRow, tmp.DayDeath.String())
		}
		dataRow = append(dataRow, strconv.Itoa(tmp.Rank))
		dataRow = append(dataRow, strconv.Itoa(tmp.GetID_dad()))
		dataRow = append(dataRow, strconv.Itoa(tmp.GetID_mom()))
		dataRow = append(dataRow, strconv.Itoa(tmp.GetID_spouse()))
		dataRow = append(dataRow, string(tmp.Gender))
		res = append(res, dataRow)
	}
	return res
}

func (pm PeopleManager) Create_saveCSV() error {
	f, err := os.OpenFile("people.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return errors.New("error in open file")
	}
	w := csv.NewWriter(f)
	data := pm.CreateData()
	for i := 0; i < pm.GetNbr(); i++ {
		dataRow := data[i]
		w.Write(dataRow)
		w.Flush()
	}
	f.Close()
	return nil
}

func StringToInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func (pm *PeopleManager) Clear() {
	pm.AllPeople = append(pm.AllPeople[:0], pm.AllPeople[pm.GetNbr():]...)
}

//read file csv
func (pm *PeopleManager) Read_CSV() error {
	fmt.Println("starting read")
	pm.Clear()
	f, err := os.Open("people.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(f)
	var list_dad []int
	var list_mom []int
	var list_spouse []int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tmp := &People{Id: int(StringToInt(record[0])), Ten: record[1], Ho: record[2], Rank: StringToInt(record[5]), Gender: record[9][0]}
		if record[3] != "nil" {
			layout := "2006-01-02T15:04:05.000Z"
			t, _ := time.Parse(layout, record[3])
			tmp.DayBirth = t
		}
		if record[4] != "nil" {
			layout := "2006-01-02T15:04:05.000Z"
			t, _ := time.Parse(layout, record[4])
			tmp.DayDeath = t
		}
		list_dad = append(list_dad, StringToInt(record[6]))
		list_mom = append(list_mom, StringToInt(record[7]))
		list_spouse = append(list_spouse, StringToInt(record[8]))
		pm.AddPeople(tmp)
	}
	for i := 0; i < pm.GetNbr(); i++ {
		if list_dad[i] != -1 {
			pm.AllPeople[i].AddDad(pm.GetPeopleByIndex(list_dad[i]))
		}
		if list_mom[i] != -1 {
			pm.AllPeople[i].AddMom(pm.GetPeopleByIndex(list_mom[i]))
		}
		if list_spouse[i] != -1 {
			pm.AllPeople[i].AddSpouse(pm.GetPeopleByIndex(list_spouse[i]))
		}
	}
	return nil
}
