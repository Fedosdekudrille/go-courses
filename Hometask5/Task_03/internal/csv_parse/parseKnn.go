package csvParse

import (
	"Task_03/pkg/knn"
	"strconv"
)

func ParseKnn(strGoods [][]string) ([]knn.Mark, []knn.DataPoint, error) {
	m := make(map[string]knn.Mark, len(strGoods))
	var d []knn.DataPoint
	mId := 0
	for _, g := range strGoods {
		s, err := strconv.ParseFloat(g[2], 64)
		if err != nil {
			return nil, nil, err
		}
		c, err := strconv.ParseFloat(g[1], 64)
		if err != nil {
			return nil, nil, err
		}
		if _, ok := m[g[0]]; !ok {
			m[g[0]] = knn.Mark{
				Id:   mId,
				Name: g[0],
			}
			mId++
		}
		d = append(d, knn.DataPoint{
			MarkId: m[g[0]].Id,
			Vector: []float64{c, s},
		})

	}
	marks := make([]knn.Mark, 0, len(m))
	for _, mark := range m {
		marks = append(marks, mark)
	}
	return marks, d, nil
}
