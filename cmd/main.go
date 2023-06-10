package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "sum", "action to apply to a certain column, default to sum operation")
	col := flag.Int64("col", 1, "col to process default to column 1")

	flag.Parse()

	if err := Run(flag.Args(), *action, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
