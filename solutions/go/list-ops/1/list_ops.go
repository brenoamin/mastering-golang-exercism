package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for i := 0; i < s.Length(); i++ {
		acc = fn(acc, s[i])
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	f := make(IntList, s.Length())
	index := 0
	for _, v := range s {
		if fn(v) {
			f[index] = v
			index++
		}
	}
	return f[:index]
}

func (s IntList) Length() int {
	length := 0
	for range s {
		length++
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	mapResult := make(IntList, s.Length())

	for i, v := range s {
		mapResult[i] = fn(v)
	}
	return mapResult
}

func (s IntList) Reverse() IntList {
	for r, l := 0, s.Length()-1; r <= l; r, l = r+1, l-1 {
		s[r], s[l] = s[l], s[r]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	newSlice := make(IntList, s.Length()+lst.Length())

	for i := 0; i < s.Length(); i++ {
		newSlice[i] = s[i]
	}

	for i := 0; i < lst.Length(); i++ {
		newSlice[s.Length()+i] = lst[i]
	}

	return newSlice
}

func (s IntList) Concat(lists []IntList) IntList {
	cs := s
	for _, list := range lists {
		cs = cs.Append(list)
	}
	return cs
}
