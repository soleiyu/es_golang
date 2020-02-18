package emsFunc

import (
//	"fmt"
	"../pictFunc"
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




