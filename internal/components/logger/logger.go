/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"github.com/mls-361/component"
)

type (
	// Logger AFAIRE.
	Logger struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Logger {
	return &Logger{
		Base: component.NewBase("logger", "logger"),
	}
}

// Dependencies AFAIRE.
func (cl *Logger) Dependencies() []string {
	return []string{
		"config",
	}
}

// Build AFAIRE.
func (cl *Logger) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
