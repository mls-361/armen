/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"os"
	"time"

	"github.com/mls-361/component"
	"github.com/mls-361/logger"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Logger AFAIRE.
	Logger struct {
		*component.Base
		components *components.Components
		logger     *logger.Master
	}
)

// New AFAIRE.
func New(components *components.Components) *Logger {
	logger := logger.New()
	components.Logger = logger

	return &Logger{
		Base:       component.NewBase("logger", "logger"),
		components: components,
		logger:     logger,
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
func (cl *Logger) Build(_ *component.Manager) error {
	app := cl.components.Application
	level := "info"

	if app.Devel() > 0 {
		level = "trace"
	}

	level, err := cl.components.Config.Data().StringWD(level, "components", "logger", "level")
	if err != nil {
		return err
	}

	if err := cl.logger.Build(app.ID(), app.Name(), level, nil); err != nil {
		return err
	}

	cl.Built()

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
