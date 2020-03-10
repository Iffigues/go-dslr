package csvs

import (
	"encoding/csv"
	"os"
)

func GetCsv(f string)(d Data, err error) {
	file, err := os.Open(f)
	if err != nil {
		return 
	}
	g := csv.NewReader(file)
	end, err := g.ReadAll()
	if err != nil {
		return
	}
	d.Names = end[0]
	d.Data = end[1:]
	return
}
