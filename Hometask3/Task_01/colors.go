package main

import "fmt"

type Color [3]byte

func (c *Color) Print() {
	fmt.Println("Red: ", c[0], "\nGreen: ", c[1], "\nBlue: ", c[2])
}

// Getters
func (c *Color) GetR() byte {
	return c[0]
}
func (c *Color) GetG() byte {
	return c[1]
}
func (c *Color) GetB() byte {
	return c[2]
}

// Setters
func (c *Color) SetR(r byte) {
	c[0] = r
}
func (c *Color) SetG(g byte) {
	c[1] = g
}
func (c *Color) SetB(b byte) {
	c[2] = b
}

func (c *Color) GetBrightness() float64 {
	return 0.2126*float64(c[0]) + 0.7152*float64(c[1]) + 0.0722*float64(c[2])
}

func maxBrightness(colors []Color) (brightest *Color) { // При colors == nil вернёт nil
	var maxBrightness, curBrightness float64
	for _, c := range colors {
		curBrightness = c.GetBrightness()
		if curBrightness > maxBrightness {
			maxBrightness = curBrightness
			brightest = &c
		}
	}
	return
}
func main() {
	var colors []Color = make([]Color, 3)
	colors[0] = Color{105, 0, 15}
	colors[1] = Color{colors[0].GetR(), 187, 3}
	colors[2] = *new(Color)
	colors[2].SetR(colors[1].GetR())
	colors[2].SetR(colors[1].GetR())
	colors[2].SetB(255)
	brightest := maxBrightness(colors)
	fmt.Println(brightest.GetBrightness())
	brightest.Print()
	fmt.Println(maxBrightness(nil))
}
