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
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components"
)

type (
	config struct {
		components *components.Components
		cfgFile    string
	}
)

func newConfig(components *components.Components) *config {
	return &config{
		components: components,
	}
}

func (c *config) defaultCfgFile() string {
	app := c.components.Application

	if cfgFile, ok := os.LookupEnv(strings.ToUpper(app.Name()) + "_CONFIG_FILE"); ok {
		return cfgFile
	}

	return filepath.Join("/etc", app.Name(), app.Name()+".yaml")
}

func (c *config) parse() error {
	app := c.components.Application
	fs := flag.NewFlagSet(app.Name(), flag.ContinueOnError)

	fs.SetOutput(os.Stdout)
	fs.Usage = func() {
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("Usage:", "armen [global options] [command [options]]")
		fmt.Println()
		fmt.Println("Global options:")
		fmt.Println()
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Commands [client mode]:")
		client.Usage(c.components)
		fmt.Println("----------------------------------------------------------------------@(°_°)@---")
		fmt.Println()
	}

	var version bool

	cfgFile := c.defaultCfgFile()

	fs.BoolVar(&version, "version", false, "print version information and quit.")
	fs.StringVar(&c.cfgFile, "config-file", cfgFile, "the YAML configuration file.")

	if err := fs.Parse(os.Args[1:]); err != nil {
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

	os.Args = append(os.Args[:1], fs.Args()...)

	return nil
}

func (c *config) load(cfg interface{}) error {
	data, err := ioutil.ReadFile(c.cfgFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}

func (c *config) build(cfg interface{}) error {
	if err := c.parse(); err != nil {
		return err
	}

	if c.components.Application.Devel() >= 2 {
		fmt.Printf("=== Config: cfgFile=%s\n", c.cfgFile) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	return c.load(cfg)
}

/*
######################################################################################################## @(°_°)@ #######
*/
