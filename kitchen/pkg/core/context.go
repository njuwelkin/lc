package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Context works as a collection of global variables for the application
// It contains the ref to logger, configuration...
type Context struct {
	// config contains the global configurations
	*config

	// out put messages to system administrator
	*audit

	// logger of kitchen system
	Log *logrus.Logger
}

func NewContext(configPath string) (*Context, error) {
	log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: false},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	config, err := newConfig(configPath)
	if err != nil {
		entry := log.WithError(err).WithField("path", configPath)
		if ResourceNotFound.Is(err) {
			entry.Warning("config file not found, use default options.")
		} else {
			entry.Error("load configuration failed")
			return nil, err
		}
	}

	ret := &Context{
		config: config,
		audit:  newAudit(nil),
		Log:    log,
	}
	err = ret.UpdateLogFileSettings()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Context) UpdateLogFileSettings() error {
	log := c.config.LogConfig
	if level, err := logrus.ParseLevel(log.Level); err == nil {
		c.Log.SetLevel(level)
	} else {
		c.Log.WithField("logLevel", log.Level).Warning("invalid log level")
		return err
	}

	if log.File != "" {
		c.Log.SetOutput(&lumberjack.Logger{
			Filename:   log.File,
			MaxSize:    log.MaxSize,
			MaxBackups: log.MaxBackups,
		})
	} else {
		// used colored log for console output
		c.Log.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02T15:04:05.999",
			FullTimestamp:   true,
		})
		c.Log.Info("log file is set to empty.  do not write log file")
	}
	return nil
}
