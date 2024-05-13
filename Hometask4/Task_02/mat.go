package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type RareMatrix[T Number] struct {
	rows, cols int
	data       map[[2]int]T
}

func NewRareMatrix[T Number](rows, cols int) *RareMatrix[T] {
	if rows <= 0 || cols <= 0 {
		panic("Неправильные координаты")
	}
	return &RareMatrix[T]{
		rows: rows,
		cols: cols,
		data: make(map[[2]int]T),
	}
}

func (m *RareMatrix[T]) Set(row, col int, value T) {
	if m == nil || row < 0 || col < 0 || row >= m.rows || col >= m.cols {
		panic("Неправильные координаты")
	}
	m.data[[2]int{row, col}] = value
}

func (m *RareMatrix[T]) Get(row, col int) T {
	if m == nil || row < 0 || col < 0 || row >= m.rows || col >= m.cols {
		panic("Неправильные координаты")
	}
	if val, ok := m.data[[2]int{row, col}]; ok {
		return val
	}
	return 0
}

type CrementInt int // Разрешает инкремент и декремент во время операции

func (i CrementInt) Inc() CrementInt {
	return i + 1
}

func (i CrementInt) Dec() CrementInt {
	return i - 1
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	m := NewRareMatrix[int](300, 300)
	m.Set(0, 0, 1)
	m.Set(150, 200, 12)
	//m.Set(300, 23, 21)
	println(m.Get(0, 0))
	println(m.Get(100, 200))
	//println(m.Get(299, 300))
	a := CrementInt(10)
	println(a.Inc().Inc().Dec() * 1000)
	cM := NewRareMatrix[CrementInt](300, 300)
	cM.Set(0, 0, a.Inc().Inc().Dec())
	println(cM.Get(0, 0))
}
