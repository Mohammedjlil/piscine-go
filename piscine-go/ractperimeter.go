package piscine

func RectPerimeter(w, h int) int {
	x := 0
	if w > 0 && h > 0 {
		a := w *2
		b := h *2
		x += a + b
	} else if w < 0 || h < 0 {
		x += -1
	

	}
	return x

}