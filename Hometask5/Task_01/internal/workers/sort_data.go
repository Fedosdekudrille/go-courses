package workers

import "sort"

type sortWSlice []Worker

func (s sortWSlice) Len() int {
	return len(s)
}

func (s sortWSlice) Less(i, j int) bool {
	if s[i].GetVacDayDur() == s[j].GetVacDayDur() { // if it can't sort by vacation days duration, then sorts by surname, then name
		for k := 0; k < len(s[i].EncSurname) && k < len(s[j].EncSurname); k++ {
			if s[i].EncSurname[k] > s[j].EncSurname[k] {
				return true
			} else if s[i].EncSurname[k] < s[j].EncSurname[k] {
				return false
			}
		}
		for k := 0; k < len(s[i].EncName) && k < len(s[j].EncName); k++ {
			if s[i].EncName[k] > s[j].EncName[k] {
				return true
			} else if s[i].EncName[k] < s[j].EncName[k] {
				return false
			}
		}
		return false
	} else if s[i].GetVacDayDur() > s[j].GetVacDayDur() {
		return true
	}
	return false
}

func (s sortWSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MapToSortedSlice(data map[string]*Worker) []Worker {
	res := make(sortWSlice, 0, len(data))
	for _, v := range data {
		res = append(res, *v)
	}
	sort.Sort(res)
	return res
}
