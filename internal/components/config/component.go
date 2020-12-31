/*
------------------------------------------------------------------------------------------------------------------------
####### config ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package config

import (
	"github.com/mls-361/component"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Config AFAIRE.
	Config struct {
		*component.Base
		config *config
	}
)

// New AFAIRE.
func New(components *components.Components) *Config {
	config := newConfig(components)
	components.Config = config

	return &Config{
		Base:   component.NewBase("config", "config"),
		config: config,
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
	if err := cc.config.build(); err != nil {
		return err
	}

	cc.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
