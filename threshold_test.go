package nagios

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestThreshold(t *testing.T) {
	is := is.New(t)

	var tests = []struct {
		threshold   string
		expectTrue  []float64
		expectFalse []float64
	}{
		{"", []float64{1, 50, 100, -1000}, nil},

		// < 0 or > 10, (outside the range of {0 .. 10})
		{"10", []float64{-1, 11}, []float64{0, 10}},

		// < 10, (outside {10 .. ∞})
		{"10:", []float64{9}, []float64{10, 50}},

		// > 10, (outside the range of {-∞ .. 10})
		{"~:10", []float64{11}, []float64{10, -50}},

		// < 10 or > 20, (outside the range of {10 .. 20})
		{"10:20", []float64{9, 21}, []float64{10, 11, 20}},

		// ≥ 10 and ≤ 20, (inside the range of {10 .. 20})
		{"@10:20", []float64{10, 15, 20}, []float64{9.5, 20.1}},
	}

	for _, test := range tests {
		thr, err := NewThreshold(test.threshold)
		is.NoErr(err)
		is.Equal(test.threshold, thr.String())

		for _, val := range test.expectTrue {
			is.True(thr.Evaluate(val))
		}

		for _, val := range test.expectFalse {
			is.False(thr.Evaluate(val))
		}

	}
}

func TestCheckThreshold(t *testing.T) {
	is := is.New(t)

	w, err := NewThreshold("80")
	is.NoErr(err)
	c, err := NewThreshold("90")
	is.NoErr(err)

	is.Equal(OK, CheckThreshold(70, w, c))
	is.Equal(WARNING, CheckThreshold(85, w, c))
	is.Equal(CRITICAL, CheckThreshold(95, w, c))
}
