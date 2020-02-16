package pictFunc

type Point struct {
	X, Y float64
}

type Line struct {
	S, D Point
}

func MkPoint(x, y float64) Point {
	var p Point

	p.X = x
	p.Y = y

	return p
}

func MkPArray(num int) []Point {
	p := make([]Point, num)

	return p
}

func MkLArray(num int) []Line {
	p := make([]Line, num)

	return p
}

func MkLineP(s, d Point) Line {
	var p Line

	p.S = s
	p.D = d

	return p
}

func MkLineF(x1, y1, x2, y2 float64) Line {
	var p Line

	p.S = MkPoint(x1, y1)
	p.D = MkPoint(x2, y2)

	return p
}
