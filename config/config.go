// config package handles setting up, parsing, and storing all configuration settings for the application.
// The main package sets up an instance of this package and all other packages can use the "GetInstance"
// method to get a reference to the config struct and all of it's values
package config

import (
	// Standard lib
	"time"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

const (
	// Environment names
	ENV_DEVELOPMENT = "dev"
	ENV_PRODUCTION  = "prod"
	ENV_TESTING     = "test"

	// Environment variable prefix
	ENV_PREFIX = "IMG_"

	// Environment variable where the path to an outside
	// configuration file is located
	CONFIG_LOCATION = ENV_PREFIX + "CONFIG"
)

var (
	config *Config
)

type (
	// Component-specific configuration

	// Struct containing configuration settings for image processing
	Images struct {
		// Max-width of the image before switching interpolators
		InterpolatorThreshold int64 `json:"interpolator-threshold" env:"IMAGE_INTERPOLATOR_THRESHOLD"`
		// Default quality all images should be output at without request overrides
		DefaultQuality int `json:"default-quality" env:"IMAGE_DEFAULT_QUALITY"`
	}

	// Struct containing configuration settings for application logging
	Log struct {
		// The formatter to use
		Formatter string `json:"formatter" env:"LOG_FORMATTER"`
		// The log level to use
		Level string `json:"level" env:"LOG_LEVEL"`
	}

	// Struct containing configuration settings for the application server
	Server struct {
		// Port the server should listen on
		Port int `json:"port" env:"SERVER_PORT"`
		// Various timeouts for the server
		Timeouts struct {
			// Timeout (in seconds) allowed for server read operations
			Read int `json:"read" env:"SERVER_READ_TIMEOUT"`
			// Timeout (in seconds) allowed for server write operations
			Write int `json:"write" env:"SERVER_WRITE_TIMEOUT"`
		} `json:"timeouts"`
	}

	// Config is a struct containing all configuration settings for the application.
	// NOTE: Only a single instance of this struct should be used throughout the application
	// so as to reference the same configuration state.
	Config struct {
		/* Top-level configuration */

		// The environment the application is running in
		Environment string `json:"environment" env:"ENVIRONMENT" usage:"dev|prod|test"`

		// Name of the application
		Name string `json:"name" env:"NAME"`

		// Start time of the application
		StartTime time.Time

		// The version of the application
		// NOTE: Used when forming and responding to HTTP requests
		Version string `json:"version" env:"VERSION"`

		/* Component-specific configuration */

		// Settings for image processing
		Images Images `json:"images"`

		// Settings for the logger
		Log Log `json:"log"`

		// Settings for the server
		Server Server `json:"server"`
	}
)

// Init sets up a configuration instance by first setting default values,
// then loading a JSON configuration file, and finally overriding and setting values
// with a corresponding environment variable. Additionally sets the application logger's various properties
func (c *Config) Init() {
	// Set up configuration defaults
	c.setDefaults()

	// Read in config file
	c.readConfigFile()

	// Set up environment overrides
	c.loadEnvironmentVariables()

	// Set logging settings
	c.setLoggerSettings()
}

// setDefault sets default configuration settings for common properties
// such as application name, environment, etc
func (c *Config) setDefaults() {
	// Top-level defaults
	c.Environment = ENV_DEVELOPMENT
	c.Name = "Img"
	c.StartTime = time.Now()
	c.Version = "v1"

	// Image defaults
	c.Images.DefaultQuality = 75
	c.Images.InterpolatorThreshold = 300

	// Logger defaults
	c.Log.Formatter = "text"
	c.Log.Level = "debug"

	// Server defaults
	c.Server.Port = 6060
	c.Server.Timeouts.Read = 30  // In seconds
	c.Server.Timeouts.Write = 30 // In seconds
}

// setLoggerSettings sets the application logger's various properties
func (c *Config) setLoggerSettings() {
	// Set logging level based on config value
	switch c.Log.Level {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	// Set logging formatter based on config value
	// NOTE: Add more cases here for custom formatters
	switch c.Log.Formatter {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}

// IsDevelopment returns a boolean indicating if the application is running
// in a development environment or not
func (c *Config) IsDevelopment() bool {
	return !c.IsProduction() && !c.IsTesting()
}

// IsProduction returns a boolean indicating if the application is running
// in a production environment or not
func (c *Config) IsProduction() bool {
	return c.Environment == ENV_PRODUCTION
}

// IsTesting returns a boolean indicating if the application is running
// in a testing environment or not
func (c *Config) IsTesting() bool {
	return c.Environment == ENV_TESTING
}

// Init creates a new config instance and initializes it
func Init() {
	// Create new config instance
	config = &Config{}

	// Initialize config
	config.Init()
}

// GetInstance returns the initialized config instance
func GetInstance() *Config {
	return config
}
