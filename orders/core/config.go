package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type shelfCapacity struct {
	Hot      int
	Cold     int
	Frozen   int
	Overflow int
}

type logConfig struct {
	File       string `yaml:"file,omitempty"`
	Level      string `yaml:"level,omitempty"`
	MaxSize    int    `yaml:"maxSize,omitempty"`
	MaxBackups int    `yaml:"maxBackups,omitempty"`
}

type config struct {
	// orders per second
	IngestRate    int           `yaml:"ingestRate"`
	ShelfCap      shelfCapacity `yaml:"shelfCapacity"`
	NumOfCouriers int           `yaml:"numOfCouriers"`

	LogConfig *logConfig `yaml:"log"`
}

const (
	defaultIngestRate      = 2
	defaultHotSelves       = 10
	defaultColdShelves     = 10
	defaultFrozenShelves   = 10
	defaultOverflowShelves = 15
	defaultNumOfCouriers   = 10
)

func NewConfig(path string) (*config, error) {
	if path == "" {
		path = defaultConfigFile
	}
	logConfig := &LogConfig{
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
		IngestRate:    defaultIngestRate,
		ShelfCap:      shelfCapacity,
		NumOfCouriers: defaultNumOfCouriers,
		LogConfig:     logConfig,
	}
	return conf, nil
}

func (c *config) load(path string) error {
	if !c.Exists(path) {
		return ResourceNotFound
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
