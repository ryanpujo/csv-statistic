package actions_test

import (
	"testing"

	"github.com/ryanpujo/csv-statistic/actions"
)

var testData = []float64{
	15000,
	10000,
	80000, 50000,
}

func TestActions(t *testing.T) {
	TestSum := 155000.0
	TestAvg := TestSum / 4.0
	testTable := map[string]struct {
		act      actions.ActFunc
		data     []float64
		expected float64
	}{
		"sum": {
			act:      actions.Sum,
			data:     testData,
			expected: TestSum,
		},
		"avg": {
			act:      actions.Avg,
			data:     testData,
			expected: TestAvg,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			actual := v.act(testData)
			if actual != float64(v.expected) {
				t.Error("failed to calculate sum/avg")
			}
		})
	}
}
