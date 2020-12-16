package main

import (
	model "github.com/phuockhanhle/familytree/model"
)

func exemple() {
	/*
		pm := model.PeopleManager{}

		p1 := pm.AddNewPerson("Phuoc Khanh", "Le", model.Male)
		p3 := pm.AddNewPerson("Kim Thai", "Le", model.Male)
		p2 := pm.AddNewPerson("Phuoc Long", "Le", model.Male)
		p4 := pm.AddNewPerson("Lien Khuong", "Tran Dinh", model.Female)
	*/
	/*
		p5 := pm.AddNewPerson("Khoi", "Tran Dinh", 'm')
		p6 := pm.AddNewPerson("Diem Chi", "Tran Dinh", 'f')
		p7 := pm.AddNewPerson("Khoi Quoc", "Tran Dinh", 'm')
		p8 := pm.AddNewPerson("Khoi Nguyen", "Tran Dinh", 'm')
		p9 := pm.AddNewPerson("Diem Phuc", "Tran Dinh", 'm')
		p10 := pm.AddNewPerson("Giang", "Tran Dinh", 'm')
		p11 := pm.AddNewPerson("Suong", "Nguyen Thi", 'f')
		p12 := pm.AddNewPerson("Ha My", "Le Phuoc", 'f')
		p13 := pm.AddNewPerson("Ba ngoai", "Nguyen", 'f')
		p14 := pm.AddNewPerson("Khuong", "Le Phuoc", 'm')

		p15 := pm.AddNewPerson("Quynh", "Tran Dinh", 'm')
		p16 := pm.AddNewPerson("Mai", "Tran Dinh", 'f')
		p17 := pm.AddNewPerson("Mien", "Tran Dinh", 'f')

		p18 := pm.AddNewPerson("Thanh Trung", "Dinh ", 'm')
		p19 := pm.AddNewPerson("Thao Nhi", "Dung", 'f')
		p20 := pm.AddNewPerson("Thanh Phuong", "Dinh", 'm')
		p21 := pm.AddNewPerson("Ong noi vis", "Dinh", 'm')
		p22 := pm.AddNewPerson("Thanh Viet", "Dinh", 'm')
		p23 := pm.AddNewPerson("Uyen My", "Dinh", 'f')

		p24 := pm.AddNewPerson("Ba co", "Nguyen", 'f')
	*/
	/*
		pm.AddNewRelation(p1.ID, p3.ID, model.ChildRole)
		pm.AddNewRelation(p3.ID, p2.ID, model.ParentRole)
		pm.AddNewRelation(p3.ID, p4.ID, model.SpouseRole)
	*/
	/*
		p1.AddDad(p3)
		p3.AddChildren(p2)
		p3.AddSpouse(p4)
		p4.AddDad(p5)
		p2.AddSpouse(p11)
		p2.AddChildren(p12)
		p5.AddChildren(p6)
		p5.AddChildren(p7)
		p5.AddChildren(p8)
		p5.AddChildren(p9)
		p5.AddChildren(p10)
		p5.AddSpouse(p13)
		p3.AddDad(p14)
		p5.AddMom(p24)
		p24.AddChildren(p15)
		p15.AddChildren(p16)
		p15.AddChildren(p17)
		p6.AddChildren(p18)
		p6.AddChildren(p19)
		p6.AddSpouse(p20)
		p20.AddDad(p21)
		p21.AddChildren(p22)
		p22.AddChildren(p23)
	*/
	/*
		model.Connect_database()
		model.Insert_person(pm.AllPeople[1])
		model.Insert_person(pm.AllPeople[0])
		model.Insert_person(pm.AllPeople[2])

		model.Insert_relation(pm.AllPeople[1], pm.AllPeople[0])
		model.Insert_relation(pm.AllPeople[1], pm.AllPeople[2])
	*/

	//model.Clear_tables()
	//pm.WriteToCSV("data")
}

func exemple_database() {
	var p1, p2, p3, p4, p5, p6, p7, p8, p9 model.Person
	p1 = model.Person{FirstName: "Phuoc Khanh", LastName: "LE", Gender: model.Male, Rank: 0, Birthday: model.StringToTime("1998-04-04T00:00:00.000Z")}
	p2 = model.Person{FirstName: "Phuoc Long", LastName: "LE", Gender: model.Male, Rank: 0, Birthday: model.StringToTime("1992-02-04T00:00:00.000Z")}
	p3 = model.Person{FirstName: "Kim Thai", LastName: "LE", Gender: model.Male, Rank: -1, Birthday: model.StringToTime("1964-04-19T00:00:00.000Z")}
	p4 = model.Person{FirstName: "Dinh Lien Khuong", LastName: "Tran", Gender: model.Female, Rank: -1, Birthday: model.StringToTime("1966-03-15T00:00:00.000Z")}
	p5 = model.Person{FirstName: "Dinh Khoi", LastName: "Tran", Gender: model.Male, Rank: -2}

	p6 = model.Person{FirstName: "Dinh Diem Chi", LastName: "Tran", Gender: model.Female, Rank: -1}
	p7 = model.Person{FirstName: "Thanh Phuong", LastName: "Dinh", Gender: model.Male, Rank: -1}
	p8 = model.Person{FirstName: "Thanh Trung", LastName: "Dinh", Gender: model.Male, Rank: 0, Birthday: model.StringToTime("1996-12-28")}
	p9 = model.Person{FirstName: "Thao Nhi", LastName: "Dinh", Gender: model.Female, Rank: 0, Birthday: model.StringToTime("2000-10-24")}

	model.Connect_database()
	model.Insert_1st_person(&p1)

	model.Addrelation(&p1, &p3, model.ChildRole)
	model.Insert_nth_person(&p1, &p3)

	model.Insert_person(&p2)
	model.Addrelation(&p3, &p2, model.ParentRole)
	model.UpdateTreeParent(&p3, &p2)

	model.Insert_person(&p4)
	model.Addrelation(&p2, &p4, model.ChildRole)
	model.UpdateTreeParent(&p2, &p4)

	model.Addrelation(&p4, &p1, model.ParentRole)
	model.MakeRelationBetweenPeopleAlreadyInDB(model.GetIdByInfo_(p1), model.GetIdByInfo_(p4), model.ChildRole)

	model.Addrelation(&p3, &p4, model.SpouseRole)
	model.MakeRelationBetweenPeopleAlreadyInDB(model.GetIdByInfo_(p3), model.GetIdByInfo_(p4), model.SpouseRole)

	model.Insert_person(&p5)
	model.Addrelation(&p5, &p4, model.ParentRole)
	model.UpdateTreeParent(&p4, &p5)

	model.Insert_person(&p6)
	model.Addrelation(&p5, &p6, model.ParentRole)
	model.UpdateTreeParent(&p5, &p6)

	model.Insert_person(&p7)
	model.Addrelation(&p6, &p7, model.SpouseRole)
	model.UpdateTreeParent(&p6, &p7)

	model.Insert_person(&p8)
	model.Addrelation(&p6, &p8, model.ParentRole)
	model.UpdateTreeParent(&p6, &p8)

	model.Insert_person(&p9)
	model.Addrelation(&p6, &p9, model.ParentRole)
	model.UpdateTreeParent(&p6, &p9)

	model.MakeRelationBetweenPeopleAlreadyInDB(model.GetIdByInfo_(p8), model.GetIdByInfo_(p7), model.ChildRole)
	model.MakeRelationBetweenPeopleAlreadyInDB(model.GetIdByInfo_(p9), model.GetIdByInfo_(p7), model.ChildRole)

}

func main() {
	//run exemple() to create file csv with people
	exemple_database()
	/*
		pm := &(model.PM)

		//verify dad and mom of all people
		pm.Read_CSV()
		for i := 0; i < pm.GetNbr(); i++ {
			if pm.AllPeople[i].Dad != nil {
				fmt.Println("dad of ", pm.AllPeople[i].Ten, " is ", pm.AllPeople[i].Dad.Ten)
			}
			if pm.AllPeople[i].Mom != nil {
				fmt.Println("mom of ", pm.AllPeople[i].Ten, " is ", pm.AllPeople[i].Mom.Ten)
			}
		}

		//verify all tree's root and its name
		model.BuildTrees("people.csv")
		for i := 0; i < len(model.TM); i++ {
			fmt.Println(model.TM[i].Filename_json)
		}

		//make file json
		model.TM[0].WriteToJson()

		roots_of_1 := model.GetRoot(pm.AllPeople[1])
		fmt.Println(roots_of_1)

		//check func SameRootByRank
		//
		j := 1
		for i := 1; i < 25; i++ {
			if i != j {
				fmt.Println(pm.GetPeopleByID(j).Ten, "is same root with ", pm.GetPeopleByID(i).Ten, "by rank ", model.RankOfSameRoot(pm.GetPeopleByID(j), pm.GetPeopleByID(i)))
			}
		}

		for i := 1; i < 25; i++ {
			if i != j {
				fmt.Println(pm.GetPeopleByID(j).Ten, "is distance with ", pm.GetPeopleByID(i).Ten, "by ", model.Distance(pm.GetPeopleByID(j), pm.GetPeopleByID(i)))
			}
		}
		//check func Get_people_in_view
		tmp := model.Get_people_in_view(pm.GetPeopleByID(j))
		fmt.Println("People in view of ", pm.GetPeopleByID(j).Ten)
		for i := 0; i < tmp.GetNbr(); i++ {
			fmt.Println(tmp.AllPeople[i].Ten)
		}
	*/

}
