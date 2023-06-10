package main

import "errors"

var (
	ErrNotNumber      = errors.New("not a number")
	ErrInvalidCol     = errors.New("invalid column")
	ErrNofiles        = errors.New("no files is provided")
	ErrInvalidActions = errors.New("invalid actions")
)
