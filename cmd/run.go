package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ryanpujo/csv-statistic/actions"
)

func Run(filenames []string, act string, col int64, out io.Writer) error {
	var actFunc actions.ActFunc

	if len(filenames) == 0 {
		return ErrNofiles
	}

	if col < 1 {
		return fmt.Errorf("%w:%d", ErrInvalidCol, col)
	}

	switch act {
	case "sum":
		actFunc = actions.Sum
	case "avg":
		actFunc = actions.Avg
	default:
		return fmt.Errorf("%w:%s", ErrInvalidActions, act)
	}

	var consolidate []float64

	for _, v := range filenames {
		f, err := os.Open(v)
		if err != nil {
			return fmt.Errorf("provided file is not available: %w:%s", ErrNofiles, err)
		}
		data, err := CsvToFloat(f, int(col))
		if err != nil {
			return fmt.Errorf("failed to parse the data: %w", err)
		}

		if err := f.Close(); err != nil {
			return fmt.Errorf("failed to close the file: %s", err)
		}

		consolidate = append(consolidate, data...)
	}

	fmt.Fprintln(out, actFunc(consolidate))
	return nil
}
