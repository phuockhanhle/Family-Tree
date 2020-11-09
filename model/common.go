package model

import (
	"math"
)

// DefaultRank is default value for rank of each person
const DefaultRank int = math.MinInt32

// IDNotFound is to represent a person that does not exist
const IDNotFound int = -1

// IDSeparator is used for string containing many IDs
const IDSeparator string = "-"

// Role is new typename for relational roles
type Role byte

// Enums for relational roles
const (
	ParentRole Role = iota
	SpouseRole
	ChildRole
	NilRole
)

// PersoRela is a map (dictionary) storing all the relations (with other people) of each person
// The key is the role of the relations, the value is the list of involved people's ids
type PersoRela map[Role][]int

// GenderType is typename for defining person's gender
type GenderType byte

// Enums for person gender
const (
	Male   GenderType = 'M'
	Female GenderType = 'F'
)

// PM is common instance of PeopleManager across the model package
var PM PeopleManager
