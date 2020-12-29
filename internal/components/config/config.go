/*
------------------------------------------------------------------------------------------------------------------------
####### config ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package config

import (
	"github.com/mls-361/component"
)

type (
	// Config AFAIRE.
	Config struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Config {
	return &Config{
		Base: component.NewBase("config", "config"),
	}
}

// Dependencies AFAIRE.
func (cc *Config) Dependencies() []string {
	return []string{
		"application",
	}
}

// Build AFAIRE.
func (cc *Config) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
