// Parallel Image Processing Benchmark

package main

import (
	"./pictFunc"
)

func main(){
	canv := pictFunc.Wcanvas(400, 600)

	canv = pictFunc.DrawLine(float64(100), float64(75), float64(200), float64(175), 5.0, canv)

	canv.Save("res.png")

}
