package knn

import "math"

type DistancePoint struct {
	MarkId   int
	Distance float64
}

func getDistance(x []float64, d DataPoint) (res float64) {
	for i := 0; i < len(x); i++ {
		res += math.Pow(x[i]-d.Vector[i], 2)
	}
	return math.Sqrt(res)
}
func NewDistancePoint(id int, dist float64) DistancePoint {
	return DistancePoint{
		MarkId:   id,
		Distance: dist,
	}
}
func replaceFarEl(x []DistancePoint, el DistancePoint) []DistancePoint {
	id := -1
	for i := len(x) - 1; i >= 0; i-- {
		if el.Distance < x[i].Distance {
			id = i
		}
	}
	if id != -1 {
		x = append(x[:id+1], x[id:len(x)-1]...)
		x[id] = NewDistancePoint(el.MarkId, el.Distance)
	}
	return x
}

func findMarkId(x []DistancePoint) int {
	total := make(map[int]int)
	for i := 0; i < len(x); i++ {
		if _, ok := total[x[i].MarkId]; ok {
			total[x[i].MarkId]++
		} else {
			total[x[i].MarkId] = 1
		}
	}
	var maxMarkId [2]int
	for id, count := range total {
		if count > maxMarkId[1] {
			maxMarkId[0] = id
			maxMarkId[1] = count
		}
	}
	return maxMarkId[0]
}
