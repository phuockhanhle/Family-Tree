package model

import (
	"time"
)

// Person is special type for each person
type Person struct {
	ID        int
	FirstName string
	LastName  string
	NickName  string
	Gender    GenderType
	Birthday  time.Time
	Deathday  time.Time
	Rank      int
	Spouse    []int
	Children  []int
	Trees     *TreeGroups
}

// GetAge returns the current age of person p
func (p Person) GetAge() int {
	return time.Now().Year() - p.Birthday.Year()
}

// UpdateRank updates the rank of person p by the rank from another person and his role to p
func (p *Person) UpdateRank(fromPerson *Person, role Role) {
	switch role {
	case ParentRole:
		p.Rank = fromPerson.Rank + 1
	case SpouseRole:
		p.Rank = fromPerson.Rank
	case ChildRole:
		p.Rank = fromPerson.Rank - 1
	}
}

/*
// AddSpouse add new spouse to list of spouse
func (p *Person) AddSpouse(s *Person) {
	if PersonAlreadyInList(s, p.Spouse) == false {
		p.Spouse = append(p.Spouse, s)
	}
}

//AddChildren add new child to list of children
func (p *Person) AddChildren(c *Person) {
	if PersonAlreadyInList(c, p.Children) == false {
		p.Children = append(p.Children, c)
	}
}
*/
//----------------------------------------------------------------------------------------------------------------

// PersonJSONForm is a convenient way to form the family trees
type PersonJSONForm struct {
	ID         int
	IDChildren []*PersonJSONForm
}

// ToJSONForm extracts useful information for PersonJSONForm from Person
/*
func (p Person) ToJSONForm() *PersonJSONForm {
	res := PersonJSONForm{ID: p.ID}
	for _, c := range p.Children {
		tmp := c.ToJSONForm()
		res.IDChildren = append(res.IDChildren, tmp)
	}
	return &res
}
*/
