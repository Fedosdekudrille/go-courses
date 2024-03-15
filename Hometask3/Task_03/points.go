package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) SetCoordinates(x, y float64) {
	p.X += x
	p.Y += y
}

func (p *Point) GetCoordinates() (x, y float64) {
	return p.X, p.Y
}

func (p *Point) Normalize(maxX, minX, maxY, minY float64) {
	if maxX == minX {
		p.X = 0
	} else {
		p.X = (p.X - minX) / (maxX - minX)
	}
	if maxY == minY {
		p.Y = 0
	} else {
		p.Y = (p.Y - minY) / (maxY - minY)
	}
}

type PointLabeled struct {
	Point
	Label string
}

func (p *PointLabeled) SetCoordinates(x, y float64) {
	p.Point.SetCoordinates(x, y)
}

func (p *PointLabeled) GetCoordinates() (x, y float64) {
	return p.Point.GetCoordinates()
}

func (p *PointLabeled) Normalize(maxX, minX, maxY, minY float64) {
	p.Point.Normalize(maxX, minX, maxY, minY)
}

type CoordinatesManager interface {
	GetCoordinates() (x, y float64)
	SetCoordinates(x, y float64)
	Normalize(maxX, minX, maxY, minY float64)
}

func getExtremeCoordinates(points []CoordinatesManager) (maxX, minX, maxY, minY float64) { //Вынес независимую логику в отдельную функцию
	if len(points) == 0 {
		return
	}
	var x, y float64
	if points[0] == nil {
		return
	}
	maxX, maxY = points[0].GetCoordinates()
	minX, minY = maxX, maxY
	for _, p := range points {
		if p == nil {
			continue
		}
		x, y = p.GetCoordinates()
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}
	return
}

func normalize(points []CoordinatesManager) {
	if len(points) < 2 {
		return // Недостаточно данных для нормализации
	}
	maxX, minX, maxY, minY := getExtremeCoordinates(points)
	for _, p := range points {
		if p == nil {
			continue
		}
		p.Normalize(maxX, minX, maxY, minY)
	}
}

func printPoints(points []CoordinatesManager) {
	for _, p := range points {
		if p == nil {
			continue
		}
		x, y := p.GetCoordinates()
		fmt.Println(x, y)
	}
}

func main() {
	var points []CoordinatesManager = make([]CoordinatesManager, 3)
	points[0] = &Point{120, -12}
	points[1] = &PointLabeled{Point{60, 16}, "1"}
	points[2] = &Point{-2, 0}
	normalize(points)
	printPoints(points)
}
