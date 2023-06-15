package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "sum", "action to apply to a certain column med|avg|sum")
	col := flag.Int64("col", 1, "col to process default to column 1")

	flag.Parse()

	/*

		use the available executable file named csv to run the program
		example:

		./path/to/csv -action [specify the action med|avg|sum] -col [specify the column] [specify csv file]

	*/

	if err := Run(flag.Args(), *action, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
