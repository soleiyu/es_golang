package pictFunc

func Wcanvas (w, h int) Pict {
	res := MkPict(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			for i := 0; i < 4; i++ {
				res.Px[x][y][i] = 255
			}
		}
	}

	return res
}
