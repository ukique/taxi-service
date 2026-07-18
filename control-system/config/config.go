package config

import "time"

type Config struct {
	Simulator SimulationConfig `yaml:"simulator"`
}
type SimulationConfig struct {
	Drivers         int           `yaml:"drivers"`
	LocationUpdates int           `yaml:"location_updates"`
	TimeOut         time.Duration `yaml:"timeout"`
}
