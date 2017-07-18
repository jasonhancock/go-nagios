package nagios

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Plugin struct {
	Name              string
	ThresholdWarning  *Threshold
	ThresholdCritical *Threshold
	flagVals          map[string]interface{}
	flagSet           *flag.FlagSet
}

func NewPlugin(name string, flag *flag.FlagSet) *Plugin {
	p := &Plugin{
		Name:     name,
		flagVals: make(map[string]interface{}),
		flagSet:  flag,
	}

	p.BoolFlag("verbose", false, "Enable verbose output")
	p.ThresholdFlag("warning", "")
	p.ThresholdFlag("critical", "")

	return p
}

func (p Plugin) Exit(code int, message string, v ...interface{}) {
	// TODO: build performance data, format the message, ec.
	fmt.Printf(message, v...)
	os.Exit(code)
}

func (p Plugin) Fatal(message interface{}) {
	fmt.Println("Fatal Error: ", message)
	os.Exit(UNKNOWN)
}

func (p Plugin) Fatalf(format string, v ...interface{}) {
	fmt.Printf("Fatal Error: "+format, v)
	os.Exit(UNKNOWN)
}

func (p Plugin) Verbose(msgs ...interface{}) {
	verbose, _ := p.OptBool("verbose")
	if verbose {
		log.Println(msgs...)
	}
}

func (p Plugin) CheckThresholds(value float64) (int, error) {
	warning, err := p.OptThreshold("warning")
	if err != nil {
		return UNKNOWN, err
	}
	critical, err := p.OptThreshold("critical")
	if err != nil {
		return UNKNOWN, err
	}

	if warning.String() != "" && critical.String() != "" {
		return CheckThreshold(value, &warning, &critical), nil
	}

	return OK, nil
}
