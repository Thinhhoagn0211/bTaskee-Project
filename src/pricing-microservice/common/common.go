package common

import (
	"encoding/json"
	"os"
)

// Configuration stores setting values
type Configuration struct {
	Port string `json:"port"`

	MgAddrs      string `json:"mgAddrs"`
	MgDbName     string `json:"mgDbName"`
	MgDbUsername string `json:"mgDbUsername"`
	MgDbPassword string `json:"mgDbPassword"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// Status Code
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000

	StatusMismatch = 10
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
