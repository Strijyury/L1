package tasks

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

// Координаты x, y
type point struct {
	x float64
	y float64
}

// Конструктор структуры Point
func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

// Вывод координаты
func (p *point) String() string {
	return fmt.Sprintf("x: %f\ny: %f", p.x, p.y)
}

// Вычисление расстояния между координатами
func Distance(p1, p2 *point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func Task24() {
	p1 := NewPoint(10, 10)
	p2 := NewPoint(25, 14)

	fmt.Println(p1)
	fmt.Println()
	fmt.Println(p2)
	fmt.Println("distance between two point is: ", Distance(p1, p2))
}
