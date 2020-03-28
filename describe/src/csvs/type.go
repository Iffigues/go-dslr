package csvs

type Data struct {
	names []string
	data  [][]string
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
