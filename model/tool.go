package model

func GetRoot(p *People) []int {
	var res []int
	if p.Dad == nil && p.Mom == nil {
		res = append(res, p.Id)
	}
	if p.Dad != nil {
		res = append(res, GetRoot(p.Dad)...)
	}
	if p.Mom != nil {
		res = append(res, GetRoot(p.Mom)...)
	}
	return res
}

func IsSameRoot(p1 *People, p2 *People) bool {
	roots_1 := GetRoot(p1)
	roots_2 := GetRoot(p2)
	for i := 0; i < len(roots_1); i++ {
		for j := 0; j < len(roots_2); j++ {
			if roots_1[i] == roots_2[j] {
				return true
			}
		}
	}
	return false
}

func Rank_of_same_root(p1 *People, p2 *People) float32 {
	if !IsSameRoot(p1, p2) {
		return -1
	} else {
		if p1.Rank == p2.Rank {
			if p1.Dad == p2.Dad || p1.Mom == p2.Mom {
				return 1
			}
			if IsSameRoot(p1.Dad, p2.Dad) {
				return Rank_of_same_root(p1.Dad, p2.Dad) + 1
			}
			if IsSameRoot(p1.Dad, p2.Mom) {
				return Rank_of_same_root(p1.Dad, p2.Mom) + 1
			}
			if IsSameRoot(p1.Mom, p2.Dad) {
				return Rank_of_same_root(p1.Mom, p2.Dad) + 1
			}
			if IsSameRoot(p1.Mom, p2.Mom) {
				return Rank_of_same_root(p1.Mom, p2.Mom) + 1
			}
		}
		if p1.Rank > p2.Rank {
			if p1.Dad == p2 || p1.Mom == p2 {
				return 0
			} else {
				if IsSameRoot(p1.Dad, p2) {
					return Rank_of_same_root(p1.Dad, p2) + 0.6
				}
				if IsSameRoot(p1.Mom, p2) {
					return Rank_of_same_root(p1.Mom, p2) + 0.6
				}
			}
		}
		if p1.Rank < p2.Rank {
			return Rank_of_same_root(p2, p1)
		}
	}
	return 0
}
