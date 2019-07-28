package nagios

import (
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFlags(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	p := NewPlugin("test", fs)

	require.NoError(t, p.StringFlag("foo", "blah", "some flag stuff"))
	require.NoError(t, p.StringFlag("foo2", "blah", "some flag stuff"))
	require.NoError(t, p.BoolFlag("bool", false, "some flag stuff"))
	require.NoError(t, p.ThresholdFlag("thresh", "0"))
	require.NoError(t, p.DurationFlag("dur", 2*time.Hour, "some duration"))
	require.NoError(t, p.DurationFlag("dur2", 2*time.Hour, "some duration"))

	fs.Parse([]string{"-foo", "bar", "-thresh", "10:20", "-dur", "1h", "-bool", "true"})

	foo, err := p.OptString("foo")
	require.NoError(t, err)
	require.Equal(t, "bar", foo)

	foo2, err := p.OptString("foo2")
	require.NoError(t, err)
	require.Equal(t, "blah", foo2)

	valBool, err := p.OptBool("bool")
	require.NoError(t, err)
	require.True(t, valBool)

	varThresh, err := p.OptThreshold("thresh")
	require.NoError(t, err)
	require.True(t, varThresh.Evaluate(9))
	require.False(t, varThresh.Evaluate(11))

	varDur, err := p.OptDuration("dur")
	require.NoError(t, err)
	require.Equal(t, 1*time.Hour, varDur)

	varDur2, err := p.OptDuration("dur2")
	require.NoError(t, err)
	require.Equal(t, 2*time.Hour, varDur2)
}
