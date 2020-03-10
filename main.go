package main

import (
	"dslr/src/csvs"
	 _ "dslr/src/data"
	"os"
	"fmt"
)

func main() {
	a, _ := csvs.GetCsv(os.Args[1])
	fmt.Println(a.Names,"\n\n\n\n" ,a.Data)
}
