package model

/*
// WIP
func GetPeopleInView(p *Person) PeopleManager {
	//
		var res PeopleManager
		for i := 0; i < PM.GetNbr(); i++ {
			if Distance(PM.AllPeople[i], p) <= 2 || RankOfSameRoot(PM.AllPeople[i], p) == 0 {
				res.AddPeople(PM.AllPeople[i])
			}
			if PM.AllPeople[i].Rank == p.Rank-1 && PM.AllPeople[i].Spouse != nil && Distance(PM.AllPeople[i].Spouse, p) == 2 {
				res.AddPeople(PM.AllPeople[i])
			}
		}
		return res
}
*/

// func To_CSV_file(p *People) error {
// 	f, err := os.OpenFile(p.Ten+"_view.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return errors.New("error in open file")
// 	}
// 	w := csv.NewWriter(f)
// 	dataRow := []string{"source", "target", "type"}
// 	w.Write(dataRow)
// 	w.Flush()
// 	for i:= 0;i<pm.GetNbr();i++ {
// 		tmp := pm.AllPeople[i]
// 		if tmp.Rank == p.Rank -3  && RankOfSameRoot( tmp,p) == 0{
// 				children := tmp.GetChildren()
// 				for j:=0;j<len(children);j++ {
// 					if
// 				}
// 		}
// 	}
// 	return nil
// }
