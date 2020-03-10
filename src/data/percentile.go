package data

import (
	"errors"
	"sort"
)

func Copy(a []float64) (b []float64) {
	for _, val := range a {
		b = append(b, val)
	}
	return
}

func Percentile(a []float64, percent float64) (percentile float64, err error) {
	length := len(a)
	c := Copy(a)
	if length == 0 {
		return 0.0, errors.New("NaN")
	}
	if length == 1 {
		return a[0], nil
	}
	if percent <= 0 || percent > 100 {
		return 0.0, errors.New("NaN")
	}
	sort.Float64s(c)
	index := (percent / 100) * float64(len(c))
	if index == float64(int64(index)) {
		i := int(index)
		percentile = c[i-1]
	} else if index > 1 {
		i := int(index)
		percentile = Mean(c[i-1 : i])

	} else {
		return 0.0, errors.New("NaN")
	}
	return percentile, nil
}
