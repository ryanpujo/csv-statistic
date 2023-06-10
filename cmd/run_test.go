package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	testTable := map[string]struct {
		act    string
		col    int
		files  []string
		assert func(t *testing.T, err error, out bytes.Buffer)
	}{
		"run sum on csv file": {
			act:   "sum",
			col:   5,
			files: []string{"../testdata/sell.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err != nil {
					t.Errorf("got an error: %s", err)
				}
				if out.String() != "31000\n" {
					t.Errorf("expected 31000, but got %s", out.String())
				}
			},
		},
		"run avg on csv file": {
			act:   "avg",
			col:   5,
			files: []string{"../testdata/sell.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				if out.String() == "10333.333333333334" {
					t.Errorf("expected 10333.333333333334, but got %s", out.String())
				}
			},
		},
		"no files provided": {
			act:   "avg",
			col:   5,
			files: []string{},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err == nil {
					t.Errorf("expected error : %s, but got nil", err)
				}
				if !errors.Is(err, ErrNofiles) {
					t.Errorf("expected error: %s but got %s", ErrNofiles, err)
				}
			},
		},
		"invalid col": {
			act:   "avg",
			col:   0,
			files: []string{"../testdata/sell.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err == nil {
					t.Errorf("expected error : %s, but got nil", err)
				}
				if !errors.Is(err, ErrInvalidCol) {
					t.Errorf("expected error: %s but got %s", ErrInvalidCol, err)
				}
			},
		},
		"files is not available": {
			act:   "avg",
			col:   5,
			files: []string{"../testdata/sold.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err == nil {
					t.Errorf("expected error : %s, but got nil", err)
				}
				if !errors.Is(err, ErrNofiles) {
					t.Errorf("expected error: %s but got %s", ErrNofiles, err)
				}
			},
		},
		"not a number": {
			act:   "avg",
			col:   2,
			files: []string{"../testdata/sell.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err == nil {
					t.Errorf("expected error : %s, but got nil", err)
				}
				if !errors.Is(err, ErrNotNumber) {
					t.Errorf("expected error: %s but got %s", ErrNotNumber, err)
				}
			},
		},
		"invalid actions": {
			act:   "som",
			col:   5,
			files: []string{"../testdata/sell.csv"},
			assert: func(t *testing.T, err error, out bytes.Buffer) {
				if err == nil {
					t.Errorf("expected error : %s, but got nil", err)
				}
				if !errors.Is(err, ErrInvalidActions) {
					t.Errorf("expected error: %s but got %s", ErrInvalidActions, err)
				}
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			var res bytes.Buffer

			err := Run(v.files, v.act, int64(v.col), &res)

			v.assert(t, err, res)
		})
	}
}
