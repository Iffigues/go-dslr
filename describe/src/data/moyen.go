package data

import (
	"math"
)

func Mean(a []float64) (t float64) {
	size := len(a)
	for _, val := range a {
		t = t + val
	}
	return t / float64(size)
}

func Min(a []float64) (b float64){
	for key, val := range a {
		if key == 0 {
			b = val
		}
		if val < b {
			b = val
		}
	}
	return
}

func Max(a []float64) (b float64){
	for _, val := range a {
		if val > b {
			b = val
		}
	}
	return
}

func Variance(a []float64) (std float64) {
	size := len(a)
        mean := Mean(a)
	 for j := 0; j < size; j++ {
                std += math.Pow(a[j]-mean, 2)
        }
	return
}

func Std(a []float64) (std float64) {
	size := len(a)
	std = Variance(a)
	return math.Sqrt(std / float64(size))
}

func Normalize(a []float64) (b []float64) {
	m := Mean(a)
	s := Std(a)
	for _, val := range a {
		b = append(b, (val-m)/s)
	}
	return
}

func Count(a [][]interface{}, aa int) (b, c float64) {
	for _, val := range  a {
		i := len(val)
		if i >= aa {
			if val[aa] != "" {
				b = b + 1.0
			} else {
				c = c + 1.0
			}
		} else {
			c = c + 1.0
		}
	}
	return
}
