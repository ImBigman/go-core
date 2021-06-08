package tapemeasure

import (
	"math"
)

// Структура для передачи координат.
type Coord struct {
	X1, Y1, X2, Y2 float64
}

// Используется для подсчета расстояния между точками.
func Distance(c Coord) float64 {
	return math.Sqrt(math.Pow(c.X2-c.X1, 2) + math.Pow(c.Y2-c.Y1, 2))
}
