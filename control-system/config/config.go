package config

import "time"

type Config struct {
	Simulator SimulationConfig `yaml:"simulator"`
	Batch     BatchConfig      `yaml:"batch"`
}

type SimulationConfig struct {
	Drivers         int           `yaml:"drivers"`
	LocationUpdates int           `yaml:"location_updates"`
	TimeOut         time.Duration `yaml:"timeout"`
}

type BatchConfig struct {
	BatchSize           int           `yaml:"batch_size"`
	WaitTimeout         time.Duration `yaml:"wait_timeout"`
	CoordinatesChanSize int           `yaml:"coordinates_chan_size"`
}
