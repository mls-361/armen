/*
------------------------------------------------------------------------------------------------------------------------
####### config ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package config

import (
	"errors"

	"github.com/mls-361/component"
	"github.com/mls-361/datamap"
	"github.com/mls-361/util"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Config AFAIRE.
	Config struct {
		*component.Base
		config *config
		data   datamap.DataMap
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
	if err := cc.config.build(&cc.data); err != nil {
		return err
	}

	cc.Built()

	return nil
}

// Data AFAIRE.
func (cc *Config) Data() datamap.DataMap {
	return cc.data
}

// Decode AFAIRE.
func (cc *Config) Decode(to interface{}, mustExist bool, keys ...string) error {
	value, err := cc.data.Retrieve(keys...)
	if err != nil {
		if errors.Is(err, datamap.ErrNotFound) {
			if mustExist {
				return err
			}

			return nil
		}

		return err
	}

	return util.DecodeData(value, to)
}

/*
######################################################################################################## @(°_°)@ #######
*/
