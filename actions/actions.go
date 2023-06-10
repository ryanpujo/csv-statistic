package actions

import "sort"

type ActFunc func(data []float64) float64

func Sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}

func Avg(data []float64) float64 {
	return Sum(data) / float64(len(data))
}

func Median(data []float64) float64 {
	sort.Float64s(data)
	if len(data)%2 == 0 {
		return (data[len(data)/2-1] + data[len(data)/2]) / 2
	}
	return data[len(data)/2]
}
