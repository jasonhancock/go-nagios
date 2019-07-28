package nagios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreshold(t *testing.T) {
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
		require.NoError(t, err)
		require.Equal(t, test.threshold, thr.String())

		for _, val := range test.expectTrue {
			require.True(t, thr.Evaluate(val))
		}

		for _, val := range test.expectFalse {
			require.False(t, thr.Evaluate(val))
		}

	}
}

func TestCheckThreshold(t *testing.T) {
	w, err := NewThreshold("80")
	require.NoError(t, err)
	c, err := NewThreshold("90")
	require.NoError(t, err)

	require.Equal(t, OK, CheckThreshold(70, w, c))
	require.Equal(t, WARNING, CheckThreshold(85, w, c))
	require.Equal(t, CRITICAL, CheckThreshold(95, w, c))
}
