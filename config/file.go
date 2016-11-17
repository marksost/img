package config

import (
	// Standard lib
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

// readConfigFile attempts to unmarshal a configuration file's contents
// from JSON into the config struct, overriding any default values set previously
func (c *Config) readConfigFile() bool {
	// Get config file contents
	contents, err := c.GetConfigFileContents()
	if err != nil {
		log.WithField("error", err.Error()).Warn("Error getting config file contents")
		return false
	}

	// Attempt to unmarshal JSON into config struct
	err = json.Unmarshal(contents, &c)
	if err != nil {
		log.WithField("error", err.Error()).Warn("Error unmarshaling JSON")
		return false
	}

	return true
}

// GetConfigFileContents attempts to load a JSON configuration file from disk and
// return it's contents if found, or an error if not
func (c *Config) GetConfigFileContents() ([]byte, error) {
	// Allow for environment-level config file location override
	file := os.Getenv(CONFIG_LOCATION)
	if file == "" {
		return nil, fmt.Errorf("No valid file path detected under environment variable %s", CONFIG_LOCATION)
	}

	// Log file location
	log.WithField("location", file).Info("Loading configuration file")

	// Attempt to load file
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.WithField("error", err.Error()).Warn("Error loading config file")

		return nil, err
	}

	return contents, nil
}
