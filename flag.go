package nagios

import (
	"github.com/pkg/errors"
)

func (p *Plugin) BoolFlag(name string, value bool, usage string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.BoolVar(p.flagVals[name].(*bool), name, value, usage)

	return nil
}

func (p Plugin) OptBool(name string) (bool, error) {

	v, ok := p.flagVals[name]
	if !ok {
		return false, errors.Errorf("option not found: %s", name)
	}

	return *v.(*bool), nil
}

func (p *Plugin) StringFlag(name string, value string, usage string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.StringVar(p.flagVals[name].(*string), name, value, usage)

	return nil
}

func (p Plugin) OptString(name string) (string, error) {

	v, ok := p.flagVals[name]
	if !ok {
		return "", errors.Errorf("option not found: %s", name)
	}

	return *v.(*string), nil
}

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

func (p *Plugin) ThresholdFlag(name string, value string) error {
	if p.flagSet.Parsed() {
		return errors.New("flag must be unparsed")
	}

	p.flagVals[name] = &value
	p.flagSet.StringVar(p.flagVals[name].(*string), name, value, name+" threshold. See https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT for format")

	return nil
}

func (p Plugin) OptThreshold(name string) (Threshold, error) {
	v, ok := p.flagVals[name]
	if !ok {
		return Threshold{}, errors.Errorf("option not found: %s", name)
	}

	t, err := NewThreshold(*v.(*string))
	if err != nil {
		return Threshold{}, err
	}

	return *t, nil
}
