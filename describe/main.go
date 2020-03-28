package main

import (
	"describe/src/csvs"
	_ "describe/src/data"
	"fmt"
	"os"
)

func main() {

	// Look if file path exists in arguments
	if len(os.Args) < 1 {
		fmt.Println("specify a file")
		os.Exit(1)
	}

	// Get CSV file contents
	if _, err := csvs.GetCsv(os.Args[1]); err == nil {
	} else {
		fmt.Println(err)
	}
}
