package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func CsvToFloat(reader io.Reader, col int) ([]float64, error) {
	cr := csv.NewReader(reader)
	cr.ReuseRecord = true
	col--

	data := make([]float64, 0, 10)

	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read file: %s", err)
		}
		if i == 0 {
			continue
		}
		if len(row) <= col {
			return nil, fmt.Errorf("%w:%d", ErrInvalidCol, col)
		}
		v, err := strconv.ParseFloat(strings.TrimSpace(row[col]), 64)
		if err != nil {
			return nil, fmt.Errorf("%w:%s", ErrNotNumber, err)
		}

		data = append(data, v)
	}
	return data, nil
}
