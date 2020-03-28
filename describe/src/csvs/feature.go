package csvs

import "strconv"

func NewFeat(v string) (g Feat) {
	g.value, g.err = strconv.ParseFloat(v, 64)
	return
}

func (r *Data) getNewFeat(i int) (f *Features) {
	f = new(Features)
	f.Name = r.names[i]
	for _, val := range r.data[i] {
		f.Data = append(f.Data, NewFeat(val))
	}
	return
}
