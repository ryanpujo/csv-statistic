package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/ryanpujo/csv-statistic/actions"
)

func Run(filenames []string, act string, col int64, out io.Writer) error {
	var actFunc actions.ActFunc
	fileCh := make(chan string)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	resCh := make(chan []float64)
	var wg sync.WaitGroup

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

	go func() {
		defer close(fileCh)
		for _, v := range filenames {
			fileCh <- v
		}
	}()

	var consolidate []float64

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range fileCh {
				f, err := os.Open(v)
				if err != nil {
					errCh <- fmt.Errorf("provided file is not available: %w:%s", ErrNofiles, err)
				}
				data, err := CsvToFloat(f, int(col))
				if err != nil {
					errCh <- fmt.Errorf("failed to parse the data: %w", err)
				}

				if err := f.Close(); err != nil {
					errCh <- fmt.Errorf("failed to close the file: %s", err)
				}

				resCh <- data
			}
		}()
		go func() {
			defer close(doneCh)
			wg.Wait()
		}()

		for {
			select {
			case data := <-resCh:
				consolidate = append(consolidate, data...)
			case err := <-errCh:
				return err
			case <-doneCh:
				_, err := fmt.Fprintln(out, actFunc(consolidate))
				return err
			}
		}
	}
	return nil
}
