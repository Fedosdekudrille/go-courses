package main

import (
	"fmt"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Clone() Cloner {
	return &Point{p.X, p.Y}
}

func (p *Point) Print() {
	fmt.Println(p.X, p.Y)
}

type Car struct {
	Mark        string
	Year        int
	EnginePower float64
}

func (c *Car) Clone() Cloner {
	return &Car{c.Mark, c.Year, c.EnginePower}
}

func (c *Car) Print() {
	fmt.Println(c.Mark, c.Year, c.EnginePower)
}

type Cloner interface {
	Clone() Cloner
}

type Printer interface {
	Print()
}

func sliceClone(c []interface{}) (cloneRes []interface{}) { // Не выделил память из-за незнания кол-ва конечных элементов
	for i := 0; i < len(c); i++ {
		switch c[i].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, bool, string:
			cloneRes = append(cloneRes, c[i]) // Если есть способ элегантнее, то я его не нашёл
		case Cloner:
			cloneRes = append(cloneRes, c[i].(Cloner).Clone())
		}
	}
	return
}

func main() {
	sliceClone(nil)
	var a []interface{} = make([]interface{}, 6)
	a[0] = strings.Builder{} // Не пройдёт
	a[1] = "Hello"
	a[2] = &Car{"BMW", 2009, 1.5}
	a[3] = a[2].(*Car).Clone()
	a[4] = &Point{1, 2}
	a[5] = nil
	cloneRes := sliceClone(a)
	for _, el := range cloneRes {
		switch el.(type) {
		case Printer:
			el.(Printer).Print()
		default:
			fmt.Println(el)
		}
	}
}
