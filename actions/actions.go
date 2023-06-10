package actions

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
