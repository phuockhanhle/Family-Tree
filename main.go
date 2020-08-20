package main

import (
	"fmt"

	"model"
)

func exemple() {
	pm := &(model.PM)

	p1 := &(model.People{Id: 1, Ten: "phuoc khanh", Ho: "le", Rank: 1, Gender: 'm'})
	p2 := &(model.People{Id: 2, Ten: "phuoc long", Ho: "le", Gender: 'm'})
	p3 := &(model.People{Id: 3, Ten: "kim thai", Ho: "le", Gender: 'm'})
	p4 := &(model.People{Id: 4, Ten: "lien khuong", Ho: "tran dinh", Gender: 'f'})
	p5 := &(model.People{Id: 5, Ten: "Khoi", Ho: "tran dinh", Gender: 'm'})
	p6 := &(model.People{Id: 6, Ten: "Diem Chi", Ho: "tran dinh", Gender: 'f'})
	p7 := &(model.People{Id: 7, Ten: "khoi quat", Ho: "tran dinh", Gender: 'm'})
	p8 := &(model.People{Id: 8, Ten: "Khoi nguyen", Ho: "tran dinh", Gender: 'm'})
	p9 := &(model.People{Id: 9, Ten: "ti", Ho: "tran dinh", Gender: 'm'})
	p10 := &(model.People{Id: 10, Ten: "giang", Ho: "tran dinh", Gender: 'm'})
	p11 := &(model.People{Id: 11, Ten: "Suong", Ho: "Nguyen thi", Gender: 'f'})
	p12 := &(model.People{Id: 12, Ten: "Ha My", Ho: "Le phuoc", Gender: 'f'})
	p13 := &(model.People{Id: 13, Ten: "Ba ngoai", Ho: "nguyen", Gender: 'f'})
	p14 := &(model.People{Id: 14, Ten: "Khuong", Ho: "Le Phuoc", Gender: 'm'})

	p15 := &(model.People{Id: 15, Ten: "Quynh", Ho: "Tran Dinh", Gender: 'm'})
	p16 := &(model.People{Id: 16, Ten: "Mai", Ho: "Tran Dinh", Gender: 'f'})
	p17 := &(model.People{Id: 17, Ten: "Mien", Ho: "Tran Dinh", Gender: 'f'})

	p18 := &(model.People{Id: 18, Ten: "Thanh Trung", Ho: "Dinh ", Gender: 'm'})
	p19 := &(model.People{Id: 19, Ten: "Thao Nhi", Ho: "Dung", Gender: 'f'})
	p20 := &(model.People{Id: 20, Ten: "Thanh Phuong", Ho: "Dinh", Gender: 'm'})
	p21 := &(model.People{Id: 21, Ten: "Ong noi vis", Ho: "Dinh", Gender: 'm'})
	p22 := &(model.People{Id: 22, Ten: "Thanh Viet", Ho: "Dinh", Gender: 'm'})
	p23 := &(model.People{Id: 23, Ten: "Uyen My", Ho: "Dinh", Gender: 'f'})

	p24 := &(model.People{Id: 24, Ten: "Ba co", Ho: "Nguyen", Gender: 'f'})

	pm.AddPeople(p1)
	pm.AddPeople(p2)
	pm.AddPeople(p3)
	pm.AddPeople(p4)
	pm.AddPeople(p5)
	pm.AddPeople(p6)
	pm.AddPeople(p7)
	pm.AddPeople(p8)
	pm.AddPeople(p9)
	pm.AddPeople(p10)
	pm.AddPeople(p11)
	pm.AddPeople(p12)
	pm.AddPeople(p13)
	pm.AddPeople(p14)
	pm.AddPeople(p15)
	pm.AddPeople(p16)
	pm.AddPeople(p17)
	pm.AddPeople(p18)
	pm.AddPeople(p19)
	pm.AddPeople(p20)
	pm.AddPeople(p21)
	pm.AddPeople(p22)
	pm.AddPeople(p23)
	pm.AddPeople(p24)

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
	pm.Create_saveCSV()
}

func main() {
	//run exemple() to create file csv with people
	//exemple()

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
	model.Build_trees("people.csv")
	for i := 0; i < len(model.TM); i++ {
		fmt.Println(model.TM[i].Filename_json)
	}

	//make file json
	model.TM[0].Savefile_Json()

	roots_of_1 := model.GetRoot(pm.AllPeople[1])
	fmt.Println(roots_of_1)

	//check func SameRootByRank
	for i := 1; i < 25; i++ {
		if i != 18 {
			fmt.Println(pm.GetPeopleByIndex(18).Ten, "is same root with ", pm.GetPeopleByIndex(i).Ten, "by rank ", model.Rank_of_same_root(pm.GetPeopleByIndex(18), pm.GetPeopleByIndex(i)))
		}
	}
}
