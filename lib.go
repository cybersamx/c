package c

import "github.com/cybersamx/e"

func Count(sample []float64) float64 {
	return float64(len(sample))
}

func Mean(sample []float64) float64 {
	sum := 0.0

	for _, v := range sample {
		sum = e.Add(sum, v)
	}

	mean := e.Divide(sum, Count(sample))

	return mean
}
