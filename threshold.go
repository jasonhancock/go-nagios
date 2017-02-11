package nagios

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var ErrorInvalidThresholdFormat error = errors.New("Invalid threshold format")

// thresholds implement the format described here: https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT
type Threshold struct {
	min         float64
	max         float64
	minInfinity bool
	maxInfinity bool
	insideRange bool
}

func NewThreshold(str string) (*Threshold, error) {
	t := &Threshold{
		min: 0,
		max: 0,
	}

	sansPrefix := strings.TrimPrefix(str, "@")
	if sansPrefix != str {
		t.insideRange = true
	}

	pieces := strings.Split(sansPrefix, ":")

	if len(pieces) == 0 || len(pieces) > 2 {
		return nil, errors.Wrap(ErrorInvalidThresholdFormat, "too many pieces")
	}

	if len(pieces) == 1 {
		// "10" means < 0 or > 10, (outside the range of {0 .. 10})
		f, err := strconv.ParseFloat(pieces[0], 64)
		if err != nil {
			return nil, errors.Wrap(ErrorInvalidThresholdFormat, "unable to parse float value")
		}
		t.max = f
		t.min = 0
	} else {
		// we have two pieces

		// "~:10" means > 10, (outside the range of {-∞ .. 10})
		if pieces[0] == "~" {
			t.minInfinity = true
		} else {
			f, err := strconv.ParseFloat(pieces[0], 64)
			if err != nil {
				return nil, errors.Wrap(ErrorInvalidThresholdFormat, "unable to parse float value")
			}
			t.min = f
		}

		if pieces[1] == "" {
			// "10:" means < 10, (outside {10 .. ∞})
			t.maxInfinity = true
		} else {
			// "10:20" means < 10 or > 20, (outside the range of {10 .. 20})
			f, err := strconv.ParseFloat(pieces[1], 64)
			if err != nil {
				return nil, errors.Wrap(ErrorInvalidThresholdFormat, "unable to parse float value")
			}
			t.max = f
		}
	}

	return t, nil
}

// Evaluate returns true if the value matches the threshold, false otherwise
func (t Threshold) Evaluate(value float64) bool {
	var inRange bool

	if t.minInfinity {
		if t.maxInfinity {
			inRange = true
		} else if value <= t.max {
			inRange = true
		}
	} else if t.maxInfinity && value >= t.min {
		inRange = true
	} else if value >= t.min && value <= t.max {
		inRange = true
	}

	if t.insideRange {
		return inRange
	}
	return !inRange
}