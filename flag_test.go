package nagios

import (
	"flag"
	"testing"

	"github.com/cheekybits/is"
)

func TestFlags(t *testing.T) {
	is := is.New(t)

	fs := flag.NewFlagSet("test", flag.PanicOnError)
	p := NewPlugin("test", fs)

	err := p.StringFlag("foo", "blah", "some flag stuff")
	is.NoErr(err)
	err = p.StringFlag("foo2", "blah", "some flag stuff")
	is.NoErr(err)
	err = p.BoolFlag("bool", false, "some flag stuff")
	is.NoErr(err)
	err = p.ThresholdFlag("thresh", "0")
	is.NoErr(err)

	fs.Parse([]string{"-foo", "bar", "-thresh", "10:20", "-bool", "true"})

	foo, err := p.OptString("foo")
	is.NoErr(err)
	is.Equal("bar", foo)

	foo2, err := p.OptString("foo2")
	is.NoErr(err)
	is.Equal("blah", foo2)

	valBool, err := p.OptBool("bool")
	is.NoErr(err)
	is.Equal(true, valBool)

	varThresh, err := p.OptThreshold("thresh")
	is.NoErr(err)
	is.True(varThresh.Evaluate(9))
	is.False(varThresh.Evaluate(11))
}
