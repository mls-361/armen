/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"os"
	"time"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/logger"
	"github.com/mls-361/minikit"
)

type (
	// Logger AFAIRE.
	Logger struct {
		*minikit.Base
		components *components.Components
		logger     *logger.Master
	}
)

// New AFAIRE.
func New(components *components.Components) *Logger {
	cl := logger.New()
	components.Logger = cl

	return &Logger{
		Base:       minikit.NewBase("logger", "logger"),
		components: components,
		logger:     cl,
	}
}

// Dependencies AFAIRE.
func (cl *Logger) Dependencies() []string {
	return []string{
		"application",
		"config",
	}
}

// Build AFAIRE.
func (cl *Logger) Build(_ *minikit.Manager) error {
	app := cl.components.Application
	level := "info"

	level, err := cl.components.Config.Data().StringWD(level, "components", "logger", "level")
	if err != nil {
		return err
	}

	if app.Debug() > 0 {
		level = "trace"
	}

	if err := cl.logger.Build(app.ID(), app.Name(), level, nil); err != nil {
		return err
	}

	cl.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"===BEGIN",
		"id", app.ID(),
		"name", app.Name(),
		"version", app.Version(),
		"builtAt", app.BuiltAt().String(),
		"pid", os.Getpid(),
	)

	return nil
}

// Close AFAIRE.
func (cl *Logger) Close() {
	cl.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"===END",
		"uptime",
		time.Since(cl.components.Application.StartedAt()).Round(time.Second).String(),
	)

	cl.logger.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
