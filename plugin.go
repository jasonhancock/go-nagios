package nagios

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Plugin stores information about a Nagios plugin.
type Plugin struct {
	Name              string
	ThresholdWarning  *Threshold
	ThresholdCritical *Threshold
	flagVals          map[string]interface{}
	flagSet           *flag.FlagSet
}

// NewPlugin initializes a new Plugin.
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

// Exit prints the plugins output message and exits with the given exit code. The
// message parameter can be a format string. fmt.Printf will be called with the
// message and the variadic arguments.
func (p Plugin) Exit(code int, message string, v ...interface{}) {
	// TODO: build performance data, format the message, ec.
	fmt.Printf(message, v...)
	os.Exit(code)
}

// Fatal aborts plugin execution, printing the message and existing with the Unknown exit status.
func (p Plugin) Fatal(message interface{}) {
	fmt.Println("Fatal Error: ", message)
	os.Exit(UNKNOWN)
}

// Fatalf aborts plugin execution, printing the message and existing with the
// Unknown exit status. The message and varadic arguments are passed on to fmt.Printf.
func (p Plugin) Fatalf(format string, v ...interface{}) {
	fmt.Printf("Fatal Error: "+format, v)
	os.Exit(UNKNOWN)
}

// Verbose logs the meesages if the plugin was called with the -verbose flag.
func (p Plugin) Verbose(msgs ...interface{}) {
	verbose, _ := p.OptBool("verbose")
	if verbose {
		log.Println(msgs...)
	}
}

// CheckThresholds checks the given value against the warning and error thresholds
// (if they were defined on the Plugin) then returns an appropriate exit code.
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
