package core

import (
	"fmt"
	//"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type shelfCapacity struct {
	// capacity of hot shelf
	Hot int `yaml:"hot"`

	// capacity of cold shelf
	Cold int `yaml:"cold"`

	// capacity of frozen shelf
	Frozen int `yaml:"frozen"`

	// capacity of overflow shelf
	Overflow int `yaml:"overflow"`
}

type logConfig struct {
	// path of log file
	File string `yaml:"file,omitempty"`

	// log level
	Level string `yaml:"level,omitempty"`

	// max log file size in MB
	MaxSize int `yaml:"maxSize,omitempty"`

	// max number of log files
	MaxBackups int `yaml:"maxBackups,omitempty"`
}

type config struct {
	// interval of ingestion in millisecond
	IngestInterval int `yaml:"ingestRate"`

	// capacity of shelves
	ShelfCap shelfCapacity `yaml:"shelfCapacity"`

	// total number of couries
	NumOfCouriers int `yaml:"numOfCouriers"`

	// log configuration
	LogConfig *logConfig `yaml:"log"`

	// debug flag
	IsDebug bool `yaml:"isDebug"`
}

const (
	defaultIngestInterval  = 500 // in millisecond
	defaultHotSelves       = 10
	defaultColdShelves     = 10
	defaultFrozenShelves   = 10
	defaultOverflowShelves = 15
	defaultNumOfCouriers   = 10

	defaultLogFile       = "./order.log"
	defaultLogLevel      = "info"
	defaultLogMaxSize    = 10 // MB
	defaultLogMaxBackups = 10

	defaultConfigFile = "./kitchen.conf"
)

func newConfig(path string) (*config, error) {
	if path == "" {
		path = defaultConfigFile
	}
	logConfig := &logConfig{
		File:       defaultLogFile,
		Level:      defaultLogLevel,
		MaxSize:    defaultLogMaxSize,
		MaxBackups: defaultLogMaxBackups,
	}
	shelfCapacity := shelfCapacity{
		Hot:      defaultHotSelves,
		Cold:     defaultColdShelves,
		Frozen:   defaultFrozenShelves,
		Overflow: defaultOverflowShelves,
	}
	conf := config{
		IngestInterval: defaultIngestInterval,
		ShelfCap:       shelfCapacity,
		NumOfCouriers:  defaultNumOfCouriers,
		LogConfig:      logConfig,
	}
	return &conf, nil
}

func (c *config) load(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return ResourceNotFound.WithField("object", "config")
		} else {
			return err
		}
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to find the config file: %w ", err)
	}
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("read config file failed: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("unmarshal conf failed: %w", err)
	}
	return nil
}
