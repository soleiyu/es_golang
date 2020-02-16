// Parallel Image Processing Benchmark

package main

import (
	"./pictFunc"
	"math"
)

func main(){
	canv := drawcircle(100)

	canv.Save("res.png")
}

func drawcircle(num int) pictFunc.Pict {

	canv := pictFunc.Wcanvas(400, 300)

	pbox := pictFunc.MkPArray(num)
	lbox := pictFunc.MkLArray(num)

	for i := 0; i < num; i++ {
		pbox[i] = pictFunc.MkPoint(
			200.0 + 100.0 * math.Cos(2.0 * math.Pi * float64(i) / float64(num)),
			150.0 + 100.0 * math.Sin(2.0 * math.Pi * float64(i) / float64(num)))
	}

	for i := 0; i < num - 1; i++ {
		lbox[i] = pictFunc.MkLineP(pbox[i], pbox[i + 1])
	}
	lbox[num - 1] = pictFunc.MkLineP(pbox[num - 1], pbox[0])

	for i := 0; i < num; i ++ {
		canv = pictFunc.DrawLineL(lbox[i], 3.0, []uint8{255, 0, 100, 200}, canv)
	}

	return canv
}
