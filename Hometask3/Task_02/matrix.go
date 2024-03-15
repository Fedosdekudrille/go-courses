package main

import "fmt"

type Matrix struct {
	rows int
	cols int
	data []float64
}

func NewMatrix(rows, cols int) *Matrix {
	if rows <= 0 || cols <= 0 {
		return &Matrix{
			rows: 1,
			cols: 1,
			data: make([]float64, 1),
		}
	}
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}
}

func (m *Matrix) Get(i, j int) float64 {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return 0
	}
	return m.data[i*m.cols+j]
}

func (m *Matrix) Set(i, j int, value float64) {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return
	}
	m.data[i*m.cols+j] = value
}

func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Println()
	}
}

func main() {
	var matrix *Matrix = NewMatrix(3, 3)
	matrix.Set(0, 0, 1)
	matrix.Set(0, 1, 2)
	matrix.Set(0, 2, 3)
	matrix.Set(1, 0, 4)
	matrix.Set(1, -1, 5)
	matrix.Set(1, 2, 6)
	matrix.Set(20, 0, 7)
	matrix.Set(2, 1, 8)
	matrix.Set(2, 2, 9)
	matrix.Print()

	fmt.Println()
	matrix = NewMatrix(0, 10)
	matrix.Print()
}
