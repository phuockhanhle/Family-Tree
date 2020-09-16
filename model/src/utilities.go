package model

import (
	"math"
	"strconv"
	"time"
)

// TimeToString converts time.Time into string. It returns nil if the time is zero
func TimeToString(t time.Time) string {
	if t.IsZero() {
		return "nil"
	}
	return t.String()
}

// StringToTime converts time string into time.Time format
func StringToTime(s string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	var t time.Time
	if s != "nil" {
		t, _ = time.Parse(layout, s)
	}
	return t
}

// StringToInt convert string into int
func StringToInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

// IsSamePerson returns if 2 pointers point to a same person
func IsSamePerson(p1 *Person, p2 *Person) bool {
	return p1.ID == p2.ID
}

// PersonAlreadyInList check if a person is already in list of people
func PersonAlreadyInList(p *Person, listP []*Person) bool {
	for _, pInList := range listP {
		if IsSamePerson(p, pInList) {
			return true
		}
	}
	return false
}

// GetRoot returns the roots of the family trees that a person belongs to
func GetRoot(p *Person) []*Person {
	var res []*Person
	if p != nil {
		if p.IsRoot() {
			res = append(res, p)
		}
		if p.Dad != nil {
			res = append(res, GetRoot(p.Dad)...)
		}
		if p.Mom != nil {
			res = append(res, GetRoot(p.Mom)...)
		}
	}
	return res
}

// IsSameRoot returns if 2 people have the same root, and the first common root of them
func IsSameRoot(p1 *Person, p2 *Person) bool {
	for _, rootP1 := range GetRoot(p1) {
		for _, rootP2 := range GetRoot(p2) {
			if IsSamePerson(rootP1, rootP2) {
				return true
			}
		}
	}
	return false
}

// FindFirstSameRootInternal calls FindFirstSameRoot and considers the result to generate its return
func FindFirstSameRootInternal(p1 *Person, p2 *Person) (*Person, *Person, *Person) {
	root, directAncestorP1, directAncestorP2 := FindFirstSameRoot(p1, p2)
	if root != nil && (directAncestorP1 == nil && directAncestorP2 == nil) {
		return root, p1, p2
	}
	return root, directAncestorP1, directAncestorP2
}

// FindFirstSameRoot returns the first (closest) common root of both people
// This is only the carrier function, calling FindFirstSameRootInternal to consider different conditions for its return
func FindFirstSameRoot(p1 *Person, p2 *Person) (*Person, *Person, *Person) {
	if IsSameRoot(p1, p2) == false {
		return nil, nil, nil
	}
	if IsSamePerson(p1, p2) == false {
		if p1.Rank == p2.Rank {
			for _, p1Parent := range []*Person{p1.Dad, p1.Mom} {
				for _, p2Parent := range []*Person{p2.Dad, p2.Mom} {
					return FindFirstSameRootInternal(p1Parent, p2Parent)
				}
			}
		}
		if p1.Rank > p2.Rank {
			for _, p1Parent := range []*Person{p1.Dad, p1.Mom} {
				return FindFirstSameRootInternal(p1Parent, p2)
			}
		}
		if p1.Rank < p2.Rank {
			return FindFirstSameRoot(p2, p1)
		}
	}
	return p1, nil, nil
}

/*
func RankOfSameRoot(p1 *Person, p2 *Person) int {
	if !IsSameRoot(p1, p2) {
		return DefaultRank
	} else {
		if p1.Rank == p2.Rank {
			if p1.Dad == p2.Dad || p1.Mom == p2.Mom {
				return 1
			}
			if IsSameRoot(p1.Dad, p2.Dad) {
				return RankOfSameRoot(p1.Dad, p2.Dad) + 1
			}
			if IsSameRoot(p1.Dad, p2.Mom) {
				return RankOfSameRoot(p1.Dad, p2.Mom) + 1
			}
			if IsSameRoot(p1.Mom, p2.Dad) {
				return RankOfSameRoot(p1.Mom, p2.Dad) + 1
			}
			if IsSameRoot(p1.Mom, p2.Mom) {
				return RankOfSameRoot(p1.Mom, p2.Mom) + 1
			}
		}
		if p1.Rank > p2.Rank {
			if p1.Dad == p2 || p1.Mom == p2 {
				return 0
			} else {
				if IsSameRoot(p1.Dad, p2) {
					return RankOfSameRoot(p1.Dad, p2)
				}
				if IsSameRoot(p1.Mom, p2) {
					return RankOfSameRoot(p1.Mom, p2)
				}
			}
		}
		if p1.Rank < p2.Rank {
			return RankOfSameRoot(p2, p1)
		}
	}
	return 0
}*/

// RankOfSameRoot computes the how much different 2 people are at their first common rank (min is chosen)
func RankOfSameRoot(p1 *Person, p2 *Person) int {
	if root, _, _ := FindFirstSameRoot(p1, p2); root != nil {
		return int(math.Min(float64(DistanceGeneration(p1, root)), float64(DistanceGeneration(p2, root))))
	}
	return DefaultRank
}

// DistanceGeneration computes the difference of 2 people's rank (abs value)
func DistanceGeneration(p1 *Person, p2 *Person) int {
	return int(math.Abs(float64(p1.Rank - p2.Rank)))
}

// Distance is a metric to help visualize people with "vertical" rank and also considering "horizontal" rank
//TODO: to be discussed
func Distance(p1 *Person, p2 *Person) int {
	return DistanceGeneration(p1, p2) + RankOfSameRoot(p1, p2)
}
