package pictFunc

import (
	"math"
)

func DrawPointP(p Point, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if queDistPoint(float64(x), float64(y), p.X, p.Y) < w {
				res.Px[x][y][0] = c[0]
				res.Px[x][y][1] = c[1]
				res.Px[x][y][2] = c[2]
				res.Px[x][y][3] = c[3]
				continue
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	return res
}

func queDistPointP(p1, p2 Point) float64 {
	return math.Sqrt((p2.X - p1.X) * (p2.X - p1.X) + 
		(p2.Y - p1.Y) * (p2.Y - p1.Y))
}

