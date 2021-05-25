package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GenderType byte

const (
	Male   GenderType = 'M'
	Female GenderType = 'F'
)

// Generate UUID
// We use V1 to guarantee uniqueness
// https://www.sohamkamani.com/uuid-versions-explained/
func GenerateID() string {
	return uuid.NewV1().String()
}

type Person struct {
	ID        string
	IDTree    string
	FirstName string
	LastName  string
	NickName  string
	Gender    GenderType
	Birthday  time.Time
	Deathday  time.Time
	HasChild  bool
	// no need to have rank anymore, since it can be replaced by 2 aspects:
	//	- use birthday to sort the list of common ancestors
	//	- use difference in length of paths to common ancestors
	//	give the difference of rank
}

// GetAge returns the current age of person p
func (p Person) GetAge() int {
	return time.Now().Year() - p.Birthday.Year()
}
