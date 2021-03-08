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
	"github.com/mls-361/minikit"
	"gopkg.in/yaml.v3"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/util"
)

type (
	// Config AFAIRE.
	Config struct {
		*minikit.Base
		components *components.Components
		flagSet    *flag.FlagSet
		cfgFile    string
		data       datamap.DataMap
	}
)

// New AFAIRE.
func New(components *components.Components) *Config {
	cc := &Config{
		Base:       minikit.NewBase("config", ""),
		components: components,
		flagSet:    flag.NewFlagSet(components.CApplication.Name(), flag.ContinueOnError),
	}

	components.CConfig = cc

	return cc
}

// Dependencies AFAIRE.
func (cc *Config) Dependencies() []string {
	return []string{
		"application",
	}
}

func (cc *Config) defaultCfgFile() string {
	app := cc.components.CApplication

	if cfgFile, ok := app.LookupEnv("CONFIG"); ok {
		return cfgFile
	}

	return filepath.Join("/etc", app.Name(), app.Name()+".yaml")
}

func (cc *Config) parse() error {
	app := cc.components.CApplication

	cc.flagSet.SetOutput(os.Stdout)
	cc.flagSet.Usage = func() {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("Usage:", app.Name(), "[global options] [command [options]]")
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
		fmt.Println("  application :", app.Name(), "v"+app.Version())
		fmt.Println("  built at    :", app.BuiltAt().String())
		fmt.Println("  copyright   :", "mls-361")
		fmt.Println("  license     :", "MIT")
		fmt.Println("-------------------------------------@(°_°)@---")
		fmt.Println()

		return client.ErrStopApp
	}

	os.Args = append(os.Args[:1], cc.flagSet.Args()...)

	return nil
}

func (cc *Config) load() error {
	data, err := ioutil.ReadFile(cc.cfgFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &cc.data)
}

// Build AFAIRE.
func (cc *Config) Build(_ *minikit.Manager) error {
	if err := cc.parse(); err != nil {
		return err
	}

	if cc.components.CApplication.Debug() > 1 {
		fmt.Printf("=== Config: cfgFile=%s\n", cc.cfgFile) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	return cc.load()
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
