package pictFunc

import (
	"fmt"
	"math"
)

func DrawLine(x1, y1, x2, y2, w float64, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {

/*
			if queDistPoint(x1, y1, float64(x), float64(y)) < w {
				res.Px[x][y][0] = 255
				res.Px[x][y][1] = 255
				res.Px[x][y][2] = 0
				res.Px[x][y][3] = 0
				continue
			} else if queDistPoint(x2, y2, float64(x), float64(y)) < w {
				res.Px[x][y][0] = 255
				res.Px[x][y][1] = 255
				res.Px[x][y][2] = 0
				res.Px[x][y][3] = 0
				continue
			}
*/

			if mmb(x1, y1, x2, y2, float64(x), float64(y)) {
				if queDistLine(x1, y1, x2, y2, float64(x), float64(y)) < w {
					fmt.Println("hoge")
					res.Px[x][y][0] = 255
					res.Px[x][y][1] = 255
					res.Px[x][y][2] = 0
					res.Px[x][y][3] = 0
					continue
				}
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	fmt.Println("DRAW LINE")
	return res
}

func mmb(x1, y1, x2, y2, x3, y3 float64) bool {
	if x3 < x1 {
		return false
	} else if x2 < x3 {
		return false
	} else if y3 < y1 {
		return false
	} else if y2 < y3 {
		return false
	}

	return true
}

func queDistPoint(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))
}

func queDistLine(x1, y1, x2, y2, x3, y3 float64) float64 {
	l := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))

	return queTriArea(x1, y1, x2, y2, x3, y3) / l
}

func queTriArea(x1, y1, x2, y2, x3, y3 float64) float64 {

	dx2 := x2 - x1
	dy2 := y2 - y1
	dx3 := x3 - x1
	dy3 := y3 - y1

	return math.Abs(dx2 * dy3 - dx3 * dy2)
}
