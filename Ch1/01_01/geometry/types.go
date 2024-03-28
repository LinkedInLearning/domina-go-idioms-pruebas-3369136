package geometry

type point struct {
	X int
	Y int
}

type Circle struct {
	Center point
	Radius int
}

func NewCircle(x int, y int, radius int) Circle {
	return Circle{
		Center: point{x, y},
		Radius: radius,
	}
}

type Triangle struct {
	V1 point
	V2 point
	V3 point
}

func NewTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) Triangle {
	return Triangle{
		V1: point{x1, y1},
		V2: point{x2, y2},
		V3: point{x3, y3},
	}
}

type Parallellogram struct {
	Min point
	Max point
}

func NewParallellogram(x1 int, y1 int, x2 int, y2 int) Parallellogram {
	return Parallellogram{
		Min: point{x1, y1},
		Max: point{x2, y2},
	}
}

func NewSquare(x1 int, y1 int, side int) Parallellogram {
	return NewParallellogram(x1, y1, x1+side, y1+side)
}

func NewRectangle(x1 int, y1 int, xSide int, ySide int) Parallellogram {
	return NewParallellogram(x1, y1, x1+xSide, y1+ySide)
}
