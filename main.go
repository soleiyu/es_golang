// Parallel Image Processing Benchmark

package main

import (
	"./pictFunc"
	"./emsFunc"
	"math"
)

func main(){
//	canv := pictFunc.Wcanvas(800, 600)
//	canv = drawcircle(200.0, 4, canv)
//	canv = drawcircle(220.0, 5, canv)

	emf := emsFunc.MkField()
	emf.AddEp(emsFunc.MkEPoint(400, 200, 30.0))
	emf.AddEp(emsFunc.MkEPoint(200, 400, -30.0))
	emf.AddEp(emsFunc.MkEPoint(400, 600, 30.0))
	emf.AddEp(emsFunc.MkEPoint(800, 200, -30.0))
	emf.AddEp(emsFunc.MkEPoint(1000, 400, 30.0))
	emf.AddEp(emsFunc.MkEPoint(800, 600, -30.0))

	canv := emf.DrawEPField(1200, 800)

	canv = emf.DrawMGMesh(10, 5, canv)

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
