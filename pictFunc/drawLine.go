package pictFunc

import (
	"fmt"
	"math"
)

func DrawLines(l []Line, w []float64, c [][]uint8, inp Pict) Pict {
	res := Wcanvas(inp.Width, inp.Height)
	res = inp

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for i := 0; i < len(l); i ++ {
				if inner3c(l[i].S.X, l[i].S.Y, l[i].D.X, l[i].D.Y, float64(x), float64(y)) {
					r := cntPointRatio(l[i], float64(x), float64(y), w[i], 1.0)
					res.Px[x][y][0] = uint8(r * float64(c[i][0]) + (1.0 - r) * float64(res.Px[x][y][0]))
					res.Px[x][y][1] = uint8(r * float64(c[i][1]) + (1.0 - r) * float64(res.Px[x][y][1]))
					res.Px[x][y][2] = uint8(r * float64(c[i][2]) + (1.0 - r) * float64(res.Px[x][y][2]))
					res.Px[x][y][3] = uint8(r * float64(c[i][3]) + (1.0 - r) * float64(res.Px[x][y][3]))
				}
			}
		}
	}

	return res
}

func DrawMMLines(l []MMLine, c [][]uint8, inp Pict) Pict {
	res := Wcanvas(inp.Width, inp.Height)
	res = inp

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for i := 0; i < len(l); i ++ {

				if float64(x) < l[i].Minx {
					continue
				} else if l[i].Maxx < float64(x) {
					continue
				} else if float64(y) < l[i].Miny {
					continue
				} else if l[i].Maxy < float64(y) {
					continue
				}


				if inner3c(l[i].S.X, l[i].S.Y, l[i].D.X, l[i].D.Y, float64(x), float64(y)) {
					r := cntPointRatioMM(l[i], float64(x), float64(y), 1.0)
					res.Px[x][y][0] = uint8(r * float64(c[i][0]) + (1.0 - r) * float64(res.Px[x][y][0]))
					res.Px[x][y][1] = uint8(r * float64(c[i][1]) + (1.0 - r) * float64(res.Px[x][y][1]))
					res.Px[x][y][2] = uint8(r * float64(c[i][2]) + (1.0 - r) * float64(res.Px[x][y][2]))
					res.Px[x][y][3] = uint8(r * float64(c[i][3]) + (1.0 - r) * float64(res.Px[x][y][3]))
				}
			}
		}
	}

	return res
}

func DrawPLP(l Line, w float64, c []uint8, inp Pict) Pict {
	ca := DrawLine5(l, w, c, inp)
	ca = DrawPointP(l.S, w, c, ca)
	ca = DrawPointP(l.D, w, c, ca)

	return ca
}

func DrawLineP(s, d Point, w float64, c []uint8, inp Pict) Pict {
	return DrawLine2(s.X, s.Y, d.X, d.Y, w, c, inp)
}

func DrawLineL(l Line, w float64, c []uint8, inp Pict) Pict {
	return DrawLine2(l.S.X, l.S.Y, l.D.X, l.D.Y, w, c, inp)
}

func pointRatio(l Line, x, y, w, ofs float64) float64 {
	dst := queDistLine(l.S.X, l.S.Y, l.D.X, l.D.Y, x, y)

	if dst < w {
		return 1.0
	} else if dst < w + ofs {
		return (w + ofs - dst) / ofs
	} else {
		return 0.0
	}
}

func inPointRatio(l Line, x, y, w, ofs float64) float64 {
	dst := queDistLine(l.S.X, l.S.Y, l.D.X, l.D.Y, x, y)

	if dst < w - ofs {
		return 1.0
	} else if dst < w {
		return (w - dst) / ofs
	} else {
		return 0.0
	}
}

func cntPointRatioMM(l MMLine, x, y, ofs float64) float64 {
	dst := queDistLine(l.S.X, l.S.Y, l.D.X, l.D.Y, x, y)

	if dst < l.W - 0.5 * ofs {
		return 1.0
	} else if dst < l.W + 0.5 * ofs {
		return (l.W - (dst - 0.5 * ofs)) / ofs
	} else {
		return 0.0
	}
}

func cntPointRatio(l Line, x, y, w, ofs float64) float64 {
	dst := queDistLine(l.S.X, l.S.Y, l.D.X, l.D.Y, x, y)

	if dst < w - 0.5 * ofs {
		return 1.0
	} else if dst < w + 0.5 * ofs {
		return (w - (dst - 0.5 * ofs)) / ofs
	} else {
		return 0.0
	}
}

func DrawLine(x1, y1, x2, y2, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if mmb(x1, y1, x2, y2, float64(x), float64(y)) {
				if queDistLine(x1, y1, x2, y2, float64(x), float64(y)) < w {
					res.Px[x][y][0] = c[0]
					res.Px[x][y][1] = c[1]
					res.Px[x][y][2] = c[2]
					res.Px[x][y][3] = c[3]
					continue
				}
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	fmt.Println("DRAW LINE")
	return res
}

func DrawLine2(x1, y1, x2, y2, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if inner3c(x1, y1, x2, y2, float64(x), float64(y)) {
				if queDistLine(x1, y1, x2, y2, float64(x), float64(y)) < w {
					res.Px[x][y][0] = c[0]
					res.Px[x][y][1] = c[1]
					res.Px[x][y][2] = c[2]
					res.Px[x][y][3] = c[3]
					continue
				}
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	fmt.Println("DRAW LINE")
	return res
}

func DrawLine3(l Line, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if inner3c(l.S.X, l.S.Y, l.D.X, l.D.Y, float64(x), float64(y)) {
				r := pointRatio(l, float64(x), float64(y), w, 1.0)
				res.Px[x][y][0] = uint8(r * float64(c[0]) + (1.0 - r) * float64(inp.Px[x][y][0]))
				res.Px[x][y][1] = uint8(r * float64(c[1]) + (1.0 - r) * float64(inp.Px[x][y][1]))
				res.Px[x][y][2] = uint8(r * float64(c[2]) + (1.0 - r) * float64(inp.Px[x][y][2]))
				res.Px[x][y][3] = uint8(r * float64(c[3]) + (1.0 - r) * float64(inp.Px[x][y][3]))
				continue
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	fmt.Println("DRAW LINE")
	return res
}

func DrawLine4(l Line, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if inner3c(l.S.X, l.S.Y, l.D.X, l.D.Y, float64(x), float64(y)) {
				r := inPointRatio(l, float64(x), float64(y), w, 1.0)
				res.Px[x][y][0] = uint8(r * float64(c[0]) + (1.0 - r) * float64(inp.Px[x][y][0]))
				res.Px[x][y][1] = uint8(r * float64(c[1]) + (1.0 - r) * float64(inp.Px[x][y][1]))
				res.Px[x][y][2] = uint8(r * float64(c[2]) + (1.0 - r) * float64(inp.Px[x][y][2]))
				res.Px[x][y][3] = uint8(r * float64(c[3]) + (1.0 - r) * float64(inp.Px[x][y][3]))
				continue
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	fmt.Println("DRAW LINE")
	return res
}

func DrawLine5(l Line, w float64, c []uint8, inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			if inner3c(l.S.X, l.S.Y, l.D.X, l.D.Y, float64(x), float64(y)) {
				r := cntPointRatio(l, float64(x), float64(y), w, 1.0)
				res.Px[x][y][0] = uint8(r * float64(c[0]) + (1.0 - r) * float64(inp.Px[x][y][0]))
				res.Px[x][y][1] = uint8(r * float64(c[1]) + (1.0 - r) * float64(inp.Px[x][y][1]))
				res.Px[x][y][2] = uint8(r * float64(c[2]) + (1.0 - r) * float64(inp.Px[x][y][2]))
				res.Px[x][y][3] = uint8(r * float64(c[3]) + (1.0 - r) * float64(inp.Px[x][y][3]))
				continue
			}
			res.Px[x][y] = inp.Px[x][y]
		}
	}

	return res
}

func DrawLine6(l MMLine, c []uint8, inp Pict) Pict {
	res := inp

	for x := 0; x < inp.Width; x++ {
		if float64(x) < l.Minx {
			continue
		} else if l.Maxx < float64(x) {
			break
		}

		for y := 0; y < inp.Height; y++ {
			if float64(y) < l.Miny {
				continue
			} else if l.Maxy < float64(y) {
				break
			}

			if inner3c(l.S.X, l.S.Y, l.D.X, l.D.Y, float64(x), float64(y)) {
				r := cntPointRatioMM(l, float64(x), float64(y), 1.0)
				res.Px[x][y][0] = uint8(r * float64(c[0]) + (1.0 - r) * float64(inp.Px[x][y][0]))
				res.Px[x][y][1] = uint8(r * float64(c[1]) + (1.0 - r) * float64(inp.Px[x][y][1]))
				res.Px[x][y][2] = uint8(r * float64(c[2]) + (1.0 - r) * float64(inp.Px[x][y][2]))
				res.Px[x][y][3] = uint8(r * float64(c[3]) + (1.0 - r) * float64(inp.Px[x][y][3]))
				continue
			}
		}
	}

	return res
}

func inner3c(sx, sy, dx, dy, cx, cy float64) bool {
	sp := MkPoint(sx, sy)
	dp := MkPoint(dx, dy)
	cp := MkPoint(cx, cy)

	return inner3p(sp, dp, cp)
}

func inner3p(sp, dp, cp Point) bool {
	if innerp2(cp, sp, dp) < 0 {
		return false
	} else if innerp2(cp, dp, sp) < 0 {
		return false
	} else {
		return true
	}
}

func innerc2(x1, y1, x2, y2, x3, y3 float64) float64 {
	p1 := MkPoint(x1, y1)
	p2 := MkPoint(x2, y2)
	p3 := MkPoint(x3, y3)

	return innerp2(p1, p2, p3)
}

func innerp2(p1, p2, p3 Point) float64 {
	v21 := MkPoint(p1.X - p2.X, p1.Y - p2.Y)
	v23 := MkPoint(p3.X - p2.X, p3.Y - p2.Y)

	return v21.X * v23.X + v21.Y * v23.Y
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
