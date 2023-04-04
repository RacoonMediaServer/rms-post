package config

import "github.com/RacoonMediaServer/rms-packages/pkg/configuration"

// Delivery specifies sending details
type Delivery struct {
	Sms struct {
		Account string
		Key     string
	}
	Smtp struct {
		Host     string
		Port     int
		User     string
		Password string
	}
}

// Configuration represents entire service configuration
type Configuration struct {
	Http     configuration.Http
	Monitor  configuration.Monitor
	Delivery Delivery
}

var config Configuration

// Load open and parses configuration file
func Load(configFilePath string) error {
	return configuration.Load(configFilePath, &config)
}

// Config returns loaded configuration
func Config() Configuration {
	return config
}
