package core

// Context works as a collection of global variables for the application
// It contains the ref to logger, configuration...
type Context struct {
	// config contains the global configurations
	*config

	// Log is a reference to the logger
	Log *logrus.Logger
}

func NewContext(configPath string) *Context {
	log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: false},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	config, err := newConfig(configPath)
	if err != nil {
		entry := log.WithError(err).WithField("path", configPath)
		if errors.Is(err, mverr.ResourceNotFound) {
			entry.Warning("config file not found, use default options.")
		} else {
			entry.Error("load configuration failed")
		}
	}

	return &Context{
		config: config,
		Log:    log,
	}
}
