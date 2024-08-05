package c

import (
	"math"
	"testing"
)

var tcs = []struct {
	m             []float64
	countExpected float64
	meanExpect    float64
}{
	{[]float64{1, 2, 3, 4}, 4, 2.5},
	{[]float64{610, 450, 160, 420, 310}, 5, 390},
}

func compare(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}

func TestCount(t *testing.T) {
	for _, tc := range tcs {
		res := Count(tc.m)
		expect := tc.countExpected

		if !compare(res, expect) {
			t.Errorf("Count: %f != %f", res, expect)
		}
	}
}

func TestMean(t *testing.T) {
	for _, tc := range tcs {
		res := Mean(tc.m)
		expect := tc.meanExpect

		if !compare(res, expect) {
			t.Errorf("Mean: %f != %f", res, expect)
		}
	}
}
