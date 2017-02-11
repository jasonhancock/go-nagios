package nagios

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestThreshold(t *testing.T) {
	is := is.New(t)

	// < 0 or > 10, (outside the range of {0 .. 10})
	thr, err := NewThreshold("10")
	is.NoErr(err)
	is.True(thr.Evaluate(-1))
	is.True(thr.Evaluate(11))
	is.False(thr.Evaluate(0))
	is.False(thr.Evaluate(10))

	// < 10, (outside {10 .. ∞})
	thr, err = NewThreshold("10:")
	is.NoErr(err)
	is.True(thr.Evaluate(9))
	is.False(thr.Evaluate(10))
	is.False(thr.Evaluate(50))

	// > 10, (outside the range of {-∞ .. 10})
	thr, err = NewThreshold("~:10")
	is.NoErr(err)
	is.True(thr.Evaluate(11))
	is.False(thr.Evaluate(10))
	is.False(thr.Evaluate(-50))

	// < 10 or > 20, (outside the range of {10 .. 20})
	thr, err = NewThreshold("10:20")
	is.NoErr(err)
	is.True(thr.Evaluate(9))
	is.True(thr.Evaluate(21))
	is.False(thr.Evaluate(10))
	is.False(thr.Evaluate(20))

	// ≥ 10 and ≤ 20, (inside the range of {10 .. 20})
	thr, err = NewThreshold("@10:20")
	is.NoErr(err)
	is.True(thr.Evaluate(10))
	is.True(thr.Evaluate(15))
	is.True(thr.Evaluate(20))
	is.False(thr.Evaluate(9.5))
	is.False(thr.Evaluate(20.1))
}
