package util

func tab_scalaire(a, b []float64) (scal float64) {
	var c []float64

	for key, val := range a {
		c = append(c, val*b[key])
	}

	for _, val := range c {
		scal = scal + val
	}
	return
}

func tab_elem_prod(a float64, b []float64) (c []float64) {

	for _, val := range b {
		c = append(c, val*a)
	}

	return
}

func tab_prod(a, b []float64) (c []float64) {

	for key, val := range a {
		c = append(c, val*b[key])
	}

	return
}

func tab_add(a, b []float64) (c []float64) {

	for key, val := range a {
		c = append(c, val+b[key])
	}

	return
}

func tab_neg(a []float64) (b []float64) {

	for _, val := range a {

		if val > 0 {
			b = append(b, val*-1)
		} else {
			b = append(b, val)
		}

	}

	return
}

func tab_abs(a []float64) (b []float64) {

	for _, val := range a {

		if val < 0 {
			b = append(b, val*-1)
		} else {
			b = append(b, val)
		}

	}

	return
}
