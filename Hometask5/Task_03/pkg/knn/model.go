package knn

import (
	"errors"
	"strconv"
)

type Mark struct {
	Id   int
	Name string
}

type DataPoint struct {
	MarkId int
	Vector []float64
}

type Knn struct {
	N     int
	marks []Mark
	data  []DataPoint //it will work even in 3 or more dimensions
}

func NewKnn(n int) (*Knn, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}
	return &Knn{
		N: n,
	}, nil
}

func (k *Knn) findMarkById(id int) (Mark, error) {
	if k.marks[id].Id == id {
		return k.marks[id], nil
	}
	for _, m := range k.marks {
		if m.Id == id {
			return m, nil
		}
	}
	return Mark{}, errors.New("marks don't content required Id")
}

func (k *Knn) Train(marks []Mark, data []DataPoint) error {
	k.marks = marks
	if len(data) == 0 {
		return errors.New("data must not be empty")
	}
	k.data = data
	for id, d := range k.data {
		if len(d.Vector) != len(k.data[0].Vector) {
			return errors.New(strconv.Itoa(id) + ": all vectors must be the same length")
		}
		if d.MarkId < 0 {
			return errors.New(strconv.Itoa(id) + ": mark Id must be > 0")
		}
		if d.MarkId > len(k.marks) {
			return errors.New(strconv.Itoa(id) + ": mark Id must be < " + strconv.Itoa(len(k.marks)))
		}
	}
	return nil
}

func (k *Knn) Predict(x []float64) (Mark, error) {
	if len(k.data) == 0 {
		return Mark{}, errors.New("model was not trained")
	}
	if len(x) == 0 {
		return Mark{}, errors.New("x must not be empty")
	}
	xLen := len(x)
	if xLen != len(k.data[0].Vector) {
		return Mark{}, errors.New("dimensions mismatch with train data")
	}
	nearestEls := make([]DistancePoint, 0, k.N)
	for _, d := range k.data {
		if len(d.Vector) != xLen {
			return Mark{}, errors.New("all vectors must be the same length")
		}
		if len(nearestEls) != k.N {
			id := len(nearestEls)
			for i := len(nearestEls) - 1; i >= 0; i-- {
				dist := getDistance(x, d)
				if i < len(nearestEls) && dist < nearestEls[i].Distance {
					id = i
				}
			}
			nearestEls = append(nearestEls[:id+1], nearestEls[id:]...)
			nearestEls[id] = NewDistancePoint(id, getDistance(x, d))
		} else {
			nearestEls = replaceFarEl(nearestEls, NewDistancePoint(d.MarkId, getDistance(x, d)))
		}
	}
	return k.findMarkById(findMarkId(nearestEls))
}
