package nagios

import (
	"errors"
	"fmt"
	"time"
)

// DurationFlag sets up a flag that handles a time.Duration value.
func (p *Plugin) DurationFlag(name string, value time.Duration, usage string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.DurationVar(p.flagVals[name].(*time.Duration), name, value, usage)

	return nil
}

// OptDuration returns the value of a given flag. A flag with that name must have
// been defined by calling DuraationFlag.
func (p Plugin) OptDuration(name string) (time.Duration, error) {
	v, ok := p.flagVals[name]
	if !ok {
		return 0, fmt.Errorf("option not found: %s", name)
	}

	return *v.(*time.Duration), nil
}

// BoolFlag sets up a flag that handles a boolean value.
func (p *Plugin) BoolFlag(name string, value bool, usage string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.BoolVar(p.flagVals[name].(*bool), name, value, usage)

	return nil
}

// OptBool returns the value of a given flag. A flag with that name must have
// been defined by calling BoolFlag.
func (p Plugin) OptBool(name string) (bool, error) {

	v, ok := p.flagVals[name]
	if !ok {
		return false, fmt.Errorf("option not found: %s", name)
	}

	return *v.(*bool), nil
}

// StringFlag sets up a flag that handles a string value.
func (p *Plugin) StringFlag(name string, value string, usage string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.StringVar(p.flagVals[name].(*string), name, value, usage)

	return nil
}

// OptString returns the value of a given flag. A flag with that name must have
// been defined by calling StringFlag.
func (p Plugin) OptString(name string) (string, error) {

	v, ok := p.flagVals[name]
	if !ok {
		return "", fmt.Errorf("option not found: %s", name)
	}

	return *v.(*string), nil
}

// OptRequiredString is like OptString, but if the value of the flag is the zero
// value due to the flag not being set, then an error is returned.
func (p Plugin) OptRequiredString(name string) string {
	val, err := p.OptString(name)
	if err != nil {
		p.Fatal(err)
	}

	if val == "" {
		p.Fatal("Required flag missing: -" + name)
	}

	return val
}

// ThresholdFlag sets up a flag that that handles a threshold value.
func (p *Plugin) ThresholdFlag(name string, value string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.StringVar(p.flagVals[name].(*string), name, value, name+" threshold. See https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT for format")

	return nil
}

// OptThreshold returns the value of a given flag. A flag with that name must have
// been defined by calling ThresholdFlag.
func (p Plugin) OptThreshold(name string) (Threshold, error) {
	v, ok := p.flagVals[name]
	if !ok {
		return Threshold{}, fmt.Errorf("option not found: %s", name)
	}

	t, err := NewThreshold(*v.(*string))
	if err != nil {
		return Threshold{}, err
	}

	return *t, nil
}
