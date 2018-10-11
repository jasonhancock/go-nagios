package nagios

import (
	"flag"
	"testing"
	"time"

	"github.com/cheekybits/is"
)

func TestFlags(t *testing.T) {
	is := is.New(t)

	fs := flag.NewFlagSet("test", flag.PanicOnError)
	p := NewPlugin("test", fs)

	is.NoErr(p.StringFlag("foo", "blah", "some flag stuff"))
	is.NoErr(p.StringFlag("foo2", "blah", "some flag stuff"))
	is.NoErr(p.BoolFlag("bool", false, "some flag stuff"))
	is.NoErr(p.ThresholdFlag("thresh", "0"))
	is.NoErr(p.DurationFlag("dur", 2*time.Hour, "some duration"))
	is.NoErr(p.DurationFlag("dur2", 2*time.Hour, "some duration"))

	fs.Parse([]string{"-foo", "bar", "-thresh", "10:20", "-dur", "1h", "-bool", "true"})

	foo, err := p.OptString("foo")
	is.NoErr(err)
	is.Equal("bar", foo)

	foo2, err := p.OptString("foo2")
	is.NoErr(err)
	is.Equal("blah", foo2)

	valBool, err := p.OptBool("bool")
	is.NoErr(err)
	is.True(valBool)

	varThresh, err := p.OptThreshold("thresh")
	is.NoErr(err)
	is.True(varThresh.Evaluate(9))
	is.False(varThresh.Evaluate(11))

	varDur, err := p.OptDuration("dur")
	is.NoErr(err)
	is.Equal(1*time.Hour, varDur)

	varDur2, err := p.OptDuration("dur2")
	is.NoErr(err)
	is.Equal(2*time.Hour, varDur2)
}
