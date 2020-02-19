package emsFunc

import (
	"fmt"
	"math"
	"../pictFunc"
)

func hello() {
	fmt.Println("Hello")
}

func UEPotential(q0, r float64) float64 {
	return q0 / r
}

func MagVctr(q EPoint, x, y float64) pictFunc.Point {
	dv := pictFunc.MkPoint(x - q.Pnt.X, y - q.Pnt.Y)
	uv := uniVctr(dv)

	return pictFunc.MkPoint(-uv.Y, uv.X)
}

func sizVctr(p pictFunc.Point) float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func uniVctr(p pictFunc.Point) pictFunc.Point {
	sv := sizVctr(p)
	return pictFunc.MkPoint(p.X / sv, p.Y / sv)
}

func malVctr(p pictFunc.Point, r float64) pictFunc.Point {
	rv := pictFunc.MkPoint(p.X * r, p.Y * r)
	return rv
}

