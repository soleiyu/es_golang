// Parallel Image Processing Benchmark

package main

import (
	"./pictFunc"
	"math"
)

func main(){
	canv := pictFunc.Wcanvas(800, 600)
	canv = drawcircle(200.0, 4, canv)
	canv = drawcircle(220.0, 5, canv)
	canv = drawcircle(240.0, 6, canv)
	canv = drawcircle(260.0, 7, canv)
	canv = drawcircle(280.0, 8, canv)
	canv = drawcircle(300.0, 9, canv)

	canv.Save("res.png")
}

func drawcircle(r float64, num int, canv pictFunc.Pict) pictFunc.Pict {

	cw := canv.Width
	ch := canv.Height

	pbox := pictFunc.MkPArray(num)
	lbox := pictFunc.MkLArray(num)

	for i := 0; i < num; i++ {
		pbox[i] = pictFunc.MkPoint(
			float64(cw) / 2.0 + r * math.Cos(2.0 * math.Pi * float64(i) / float64(num)),
			float64(ch) / 2.0 + r * math.Sin(2.0 * math.Pi * float64(i) / float64(num)))
	}

	for i := 0; i < num - 1; i++ {
		lbox[i] = pictFunc.MkLineP(pbox[i], pbox[i + 1])
	}
	lbox[num - 1] = pictFunc.MkLineP(pbox[num - 1], pbox[0])

	for i := 0; i < num; i ++ {
		canv = pictFunc.DrawPLP(lbox[i], 1.0,
			[]uint8{255, uint8(255 - (255 / num) * i),
				0, uint8(255 / num * i)}, canv)
	}

	return canv
}
