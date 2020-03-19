package main

import (
	"describe/src/csvs"
	 _ "describe/src/data"
	"os"
	"fmt"
)

func main() {

	// Look if file path exists in arguments
	if len(os.Args) > 1 {
		fmt.Println("specify a file")
		os.Exit(1)
	}

	// Get CSV file contents
	if a, err := csvs.GetCsv(os.Args[1]); err == nil {
		fmt.Println(a.Names,"\n\n\n\n" ,a.Data)
	} else {
		fmt.Println(err)
	}
}
