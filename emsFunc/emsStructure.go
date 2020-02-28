package emsFunc

import (
	"fmt"
	"../pictFunc"
	"math"
)

type EPoint struct {
	Pnt pictFunc.Point
	Pwr float64
}

type Field struct {
	//Electric Point
	Eps []EPoint
}

//Constractor
func MkEPoint(x, y, pwr float64) EPoint {
	var p EPoint
	p.Pnt = pictFunc.MkPoint(x, y)
	p.Pwr = pwr

	return p
}

//Constractor
func MkField() Field {
	var p Field
	p.Eps = make([]EPoint, 0)

	return p
}

func (this *Field) AddEp(inp EPoint) {
	this.Eps = append(this.Eps, inp)
}

func (this *Field) CalcMGVLS(inp pictFunc.Point, val float64) pictFunc.Line {
	vp := pictFunc.MkPoint(0.0, 0.0)

	for i := 0; i < len(this.Eps); i++ {
		cp := MagVctr(this.Eps[i], inp.X, inp.Y)
		d := fdist(this.Eps[i], inp.X, inp.Y)
		uep := UEPotential(this.Eps[i].Pwr, d)

		vp.X += uep * cp.X
		vp.Y += uep * cp.Y
	}

	uv := uniVctr(vp)
	vp.X = val * uv.X
	vp.Y = val * uv.Y

	rl := pictFunc.MkLineP(inp, pictFunc.MkPoint(inp.X + vp.X, inp.Y + vp.Y))

	return rl
}

func (this *Field) CalcMGVL(inp pictFunc.Point, lim float64) pictFunc.Line {
	vp := pictFunc.MkPoint(0.0, 0.0)

	for i := 0; i < len(this.Eps); i++ {
		cp := MagVctr(this.Eps[i], inp.X, inp.Y)
		d := fdist(this.Eps[i], inp.X, inp.Y)
		uep := UEPotential(this.Eps[i].Pwr, d)

		vp.X += uep * cp.X
		vp.Y += uep * cp.Y
	}

	td := fpdist(vp, 0.0, 0.0)

	if lim < td {
		uv := uniVctr(vp)
		vp.X = lim * uv.X
		vp.Y = lim * uv.Y
	}

	rl := pictFunc.MkLineP(inp, pictFunc.MkPoint(inp.X + vp.X, inp.Y + vp.Y))

	return rl
}

func (this *Field) CalcMGV(inp pictFunc.Point) pictFunc.Point {
	x := 0.0
	y := 0.0

	for i := 0; i < len(this.Eps); i++ {
		cp := MagVctr(this.Eps[i], inp.X, inp.Y)
		d := fdist(this.Eps[i], inp.X, inp.Y)

		x += d * cp.X
		y += d * cp.Y
	}

	return pictFunc.MkPoint(x, y)
}

func (this *Field) DrawMGMesh2(pitch, ofs float64, inp pictFunc.Pict) pictFunc.Pict {
	res := inp

	lb := make([]pictFunc.Line, 0)
	cb := make([][]uint8, 0)
	wb := make([]float64, 0)

	for x := ofs; x < float64(inp.Width); x += pitch {
		for y := ofs; y < float64(inp.Height); y += pitch {
			epos := MkEPoint(x, y, 1)
			mgl := this.CalcMGVLS(epos.Pnt, 7)
			lb = append(lb, mgl)
			cb = append(cb, []uint8{255, 0, 255, 0})
			wb = append(wb, 1.0)

			fmt.Println(x, y)
		}
	}

	res = pictFunc.DrawLines(lb, wb, cb, inp)

	return res
}

func (this *Field) DrawMGMesh(pitch, ofs float64, inp pictFunc.Pict) pictFunc.Pict {
	res := inp

	for x := ofs; x < float64(inp.Width); x += pitch {
		for y := ofs; y < float64(inp.Height); y += pitch {
			epos := MkEPoint(x, y, 1)
			mgl := this.CalcMGVLS(epos.Pnt, 7)
			res = pictFunc.DrawPLP(mgl, 1.0, []uint8{255, 0, 255, 0}, res)

			fmt.Println(x, y)
		}
	}

	return res
}

func (this *Field) DrawEPField(w, h int) pictFunc.Pict {
	canv := pictFunc.MkPict(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			vsum := 0.0

			for i := 0; i < len(this.Eps); i++ {
				vsum += UEPotential(this.Eps[i].Pwr,
					dist(this.Eps[i], x, y))
			}

			if 1.0 < vsum {
				vsum = 1.0
			} else if vsum < -1.0 {
				vsum = -1.0
			}

			if vsum < 0 {
				canv.Px[x][y][0] = 255
				canv.Px[x][y][1] = 255 - uint8(-float64(255) * vsum)
				canv.Px[x][y][2] = 255 - uint8(-float64(255) * vsum)
				canv.Px[x][y][3] = 255
			} else {
				canv.Px[x][y][0] = 255
				canv.Px[x][y][1] = 255
				canv.Px[x][y][2] = 255 - uint8(float64(255) * vsum)
				canv.Px[x][y][3] = 255 - uint8(float64(255) * vsum)
			}
		}
	}

	return canv
}

func dist(ep EPoint, x, y int) float64 {
	return math.Sqrt(math.Pow(float64(x) - ep.Pnt.X, 2.0) +
		math.Pow(float64(y) - ep.Pnt.Y, 2.0))
}

func fdist(ep EPoint, x, y float64) float64 {
	return math.Sqrt(math.Pow(x - ep.Pnt.X, 2.0) +
		math.Pow(y - ep.Pnt.Y, 2.0))
}

func fpdist(p pictFunc.Point, x, y float64) float64 {
	return math.Sqrt(math.Pow(x - p.X, 2.0) +
		math.Pow(y - p.Y, 2.0))
}




