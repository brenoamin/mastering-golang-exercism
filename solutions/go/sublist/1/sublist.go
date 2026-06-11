package sublist

import "slices"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	l1Length := len(l1)
	l2Length := len(l2)

	if l1Length == l2Length {
		if slices.Equal(l1, l2) || l1Length == 0 && l2Length == 0 {
			return "equal"
		} else {
			return "unequal"
		}
	} else {
		if l1Length == 0 {
			if l2Length != 0 {
				return "sublist"
			}
		} else {
			if l2Length == 0 {
				return "superlist"
			} else {
				if l1Length > l2Length {
					if isSublist(l2, l1) {
						return "superlist"
					} else {
						return "unequal"
					}
				} else {
					if isSublist(l1, l2) {
						return "sublist"
					} else {
						return "unequal"
					}
				}
			}
		}
	}
	return "unequal"
}

func isSublist(a, b []int) bool {
	var result int
	for l, r := 0, len(a)-1; r <= len(b)-1; l, r = l+1, r+1 {
		for i, v := range b[l : r+1] {
			if a[i] == v {
				result++
			} else {
				result = 0
				break
			}
		}
		if result == len(a) {
			break
		}
	}
	return result == len(a)
}
