package main

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

var testData = `
no,name,price,units sold,total_revenue
1,macbook,3000,5,15000
2,iphone,1200,10,12000
3,pixel,800,30,4000
`

func TestCsv2Float(t *testing.T) {
	col5 := []float64{15000.0, 12000.0, 4000.0}
	testTable := map[string]struct {
		col    int
		r      io.Reader
		assert func(t *testing.T, actual []float64, err error)
	}{
		"col5": {
			col: 5,
			r:   bytes.NewReader([]byte(testData)),
			assert: func(t *testing.T, actual []float64, err error) {
				if err != nil {
					t.Errorf("got an error")
				}
				if len(actual) != len(col5) {
					t.Errorf("actual result is not the same as expected")
				}
			},
		},
		"not anumber": {
			col: 2,
			r:   bytes.NewReader([]byte(testData)),
			assert: func(t *testing.T, actual []float64, err error) {
				if actual != nil {
					t.Errorf("actual is not nil")
				}
				if err == nil {
					t.Errorf("expected an error but got nil")
				}
				if !errors.Is(err, ErrNotNumber) {
					t.Errorf("expected error: %s, but got %s", ErrNotNumber, err)
				}
			},
		},
		"invalid col": {
			col: 6,
			r:   bytes.NewReader([]byte(testData)),
			assert: func(t *testing.T, actual []float64, err error) {
				if actual != nil {
					t.Errorf("actual is not nil")
				}
				if err == nil {
					t.Errorf("expected an error but got nil")
				}
				if !errors.Is(err, ErrInvalidCol) {
					t.Errorf("expected error: %s, but got %s", ErrInvalidCol, err)
				}
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			actual, err := CsvToFloat(v.r, v.col)

			v.assert(t, actual, err)
		})
	}
}
