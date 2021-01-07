/*
------------------------------------------------------------------------------------------------------------------------
####### config ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package config

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Config AFAIRE.
	Config struct {
		*minikit.Base
		config *cConfig
	}
)

// New AFAIRE.
func New(components *components.Components) *Config {
	config := newCConfig(components)
	components.Config = config

	return &Config{
		Base:   minikit.NewBase("config", "config"),
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
func (cc *Config) Build(_ *minikit.Manager) error {
	return cc.config.build()
}

/*
######################################################################################################## @(°_°)@ #######
*/
