// Parallel Image Processing Benchmark

package main

import (
	"./pictFunc"
	"math"
)

func main(){
	canv := drawcircle(250.0, 7)

	canv.Save("res.png")
}

func drawcircle(r float64, num int) pictFunc.Pict {

	cw := 800
	ch := 600

	canv := pictFunc.Wcanvas(cw, ch)

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
		canv = pictFunc.DrawPL(lbox[i], 3.0, 
			[]uint8{255, uint8(250 - 25 * i), 
				0, uint8(25 * i)}, canv)
	}

	return canv
}
