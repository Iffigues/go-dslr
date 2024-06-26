package csvs

import "strconv"

func NewFeat(v, f string) (g Feat) {
	if f == "" {

	}
	g.value, g.err = strconv.ParseFloat(v, 64)
	return
}

func (r *Data) GetNewFeat(i int) (f *Features) {
	f = new(Features)
	var d []string
	f.Name = r.names[i]
	for _, val := range r.data {
		d = append(d, val[i])
	}
	for _, val := range d {
		f.Data = append(f.Data, NewFeat(val, f.Name))
	}
	return
}
