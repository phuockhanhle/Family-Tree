package model

import (
	"errors"
	"time"
)

type error interface {
	Error() string
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

var PM PeopleManager

type People struct {
	Id       int
	Ten      string
	Ho       string
	DayBirth time.Time
	DayDeath time.Time
	Rank     int
	Dad      *People
	Mom      *People
	Spouse   *People
	Gender   byte
}

// func (PM PeopleManager) init() {
// 	PM.nbr = 0
// }

//var peopleRegex = regexp.MustCompile("^\\x*S")

func (p People) GetID_dad() int {
	if p.Dad == nil {
		return -1
	}
	return (*p.Dad).Id
}

func (p People) GetID_mom() int {
	if p.Mom == nil {
		return -1
	}
	return (*p.Mom).Id
}

func (p People) GetID_spouse() int {
	if p.Spouse == nil {
		return -1
	}
	return (*p.Spouse).Id
}

func (p People) GetAge() int {
	return time.Now().Year() - p.DayBirth.Year()
}

func (parent *People) GetChildren() []*People {
	var res []*People
	for i := 0; i < PM.GetNbr(); i++ {
		if PM.AllPeople[i].Dad == parent || PM.AllPeople[i].Mom == parent {
			res = append(res, PM.AllPeople[i])
		}
	}
	return res
}

func (p *People) AddDad(d *People) (int, error) {
	if p.Dad != nil {
		return 0, errors.New("already have dad")
	}
	if d.Spouse == nil {
		if (*p).Mom != nil {
			(*d).Spouse = (*p).Mom
		}
	}
	if d.Spouse != nil {
		if (*p).Mom == nil {
			(*p).Mom = (*d).Spouse
		}
		if (*p).Mom != (*d).Spouse {
			return 0, errors.New("mom and wife of dad is not same")
		}
	}
	d.Rank = p.Rank - 1
	p.Dad = d
	return 1, nil
}

func (p *People) AddMom(m *People) (int, error) {
	if p.Mom != nil {
		return 0, errors.New("already have mom")
	}
	if (*m).Spouse == nil {
		if (*p).Dad != nil {
			(*m).Spouse = (*p).Dad
		}
	}
	if (*m).Spouse != nil {
		if (*p).Dad == nil {
			(*p).Dad = (*m).Spouse
		}
		if (*p).Dad != (*m).Spouse {
			return 0, errors.New("dad and husband of mom is not same")
		}
	}
	m.Rank = p.Rank - 1
	p.Mom = m
	return 1, nil
}

func (p *People) AddSpouse(sp *People) error {
	if (*p).Spouse != nil {
		return errors.New("already has spouse")
	}
	sp.Rank = p.Rank
	children := p.GetChildren()
	for i := 0; i < len(children); i++ {
		if sp.Gender == 'f' {
			children[i].AddMom(sp)
		}
		if sp.Gender == 'm' {
			children[i].AddDad(sp)
		}
	}
	(*p).Spouse = sp
	return nil
}

func (p *People) AddChildren(c *People) {
	c.Rank = p.Rank + 1
	if p.Gender == 'f' {
		c.AddMom(p)
	}
	if p.Gender == 'm' {
		c.AddDad(p)
	}
}
