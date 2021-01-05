/*
------------------------------------------------------------------------------------------------------------------------
####### config ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mls-361/datamap"
	"gopkg.in/yaml.v3"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/util"
)

type (
	cConfig struct {
		components *components.Components
		flagSet    *flag.FlagSet
		cfgFile    string
		data       datamap.DataMap
	}
)

func newCConfig(components *components.Components) *cConfig {
	return &cConfig{
		components: components,
		flagSet:    flag.NewFlagSet(components.Application.Name(), flag.ContinueOnError),
	}
}

func (cc *cConfig) defaultCfgFile() string {
	app := cc.components.Application

	if cfgFile, ok := app.LookupEnv("CONFIG"); ok {
		return cfgFile
	}

	return filepath.Join("/etc", app.Name(), app.Name()+".yaml")
}

func (cc *cConfig) parse() error {
	app := cc.components.Application

	cc.flagSet.SetOutput(os.Stdout)
	cc.flagSet.Usage = func() {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("Usage:", "armen [global options] [command [options]]")
		fmt.Println()
		fmt.Println("Global options:")
		fmt.Println()
		cc.flagSet.PrintDefaults()
		fmt.Println()
		fmt.Println("Commands [client mode]:")
		client.Usage(cc.components)
		fmt.Println("----------------------------------------------------------------------@(°_°)@---")
		fmt.Println()
	}

	var version bool

	cfgFile := cc.defaultCfgFile()

	cc.flagSet.BoolVar(&version, "version", false, "print version information and quit.")
	cc.flagSet.StringVar(&cc.cfgFile, "config", cfgFile, "the YAML configuration file.")

	if err := cc.flagSet.Parse(os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return client.ErrStopApp
		}

		return err
	}

	if version {
		fmt.Println()
		fmt.Println("-----------------------------------------------")
		fmt.Println("  armen     :", "v"+app.Version())
		fmt.Println("  built at  :", app.BuiltAt().String())
		fmt.Println("  copyright :", "mls-361")
		fmt.Println("  license   :", "MIT")
		fmt.Println("-------------------------------------@(°_°)@---")
		fmt.Println()

		return client.ErrStopApp
	}

	os.Args = append(os.Args[:1], cc.flagSet.Args()...)

	return nil
}

func (cc *cConfig) load() error {
	data, err := ioutil.ReadFile(cc.cfgFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &cc.data)
}

func (cc *cConfig) build() error {
	if err := cc.parse(); err != nil {
		return err
	}

	if cc.components.Application.Debug() > 1 {
		fmt.Printf("=== Config: cfgFile=%s\n", cc.cfgFile) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	return cc.load()
}

// Data AFAIRE.
func (cc *cConfig) Data() datamap.DataMap {
	return cc.data
}

// Decode AFAIRE.
func (cc *cConfig) Decode(to interface{}, mustExist bool, keys ...string) error {
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
