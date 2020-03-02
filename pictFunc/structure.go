package pictFunc

type Point struct {
	X, Y float64
}

type MMLine struct {
	S, D Point
	W, Minx, Miny, Maxx, Maxy float64
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

func MkMMLineL(w float64, l Line) MMLine {
	return MkMMLineF(l.S.X, l.S.Y, l.D.X, l.D.Y, w)
}

func MkMMLineF(x1, y1, x2, y2, w float64) MMLine {
	var p MMLine

	p.S = MkPoint(x1, y1)
	p.D = MkPoint(x2, y2)
	p.W = w

	if x1 < x2 {
		p.Minx = x1 - w
		p.Maxx = x2 + w
	} else {
		p.Minx = x2 - w
		p.Maxx = x1 + w
	}

	if y1 < y2 {
		p.Miny = y1 - w
		p.Maxy = y2 + w
	} else {
		p.Miny = y2 - w
		p.Maxy = y1 + w
	}

	return p
}
