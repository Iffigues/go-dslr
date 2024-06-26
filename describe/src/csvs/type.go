package csvs

type oneHot struct {
	name string
	one  [4]float64
}

type Data struct {
	names []string
	data  [][]string
	feat  []*Features
}

type Feat struct {
	err   error
	value float64
}

type Features struct {
	Name string
	Data []Feat
}

func (v *Data) GetValues(f string) (b []string) {
	for key, val := range v.names {
		if val == f {
			for _, val := range v.data {
				b = append(b, val[key])
			}
			return
		}
	}
	return
}
