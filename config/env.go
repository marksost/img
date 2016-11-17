package config

import (
	// Standard lib
	"flag"
	"os"
	"reflect"
	"strconv"
	"strings"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

// loadEnvironmentVariables attempts to load environment variables matching
// the config struct's env tags, overriding any default or file-based values set previously
func (c *Config) loadEnvironmentVariables() {
	// Log environment variable read operation
	log.Info("Reading environment variables into config")

	// Start handling fields
	c.handleFields(reflect.ValueOf(c))
}

// handleFields loops through a reflected value's fields by their "kind", checks for a corresponding environment
// variable and, if found, sets it both on the config and as a flag (when allowed)
func (c *Config) handleFields(temp reflect.Value) {
	// Reflect value indirectly so as to loop through it's fields
	value := reflect.Indirect(temp)

	// Loop through fields
	for i := 0; i < value.NumField(); i++ {
		// Store kind, env tag value, flag name, and OS value
		kind := value.Field(i).Kind()
		tag := ENV_PREFIX + value.Type().Field(i).Tag.Get("env")
		flagName := c.formFlagName(tag)
		// NOTE: Enforces upper-case env variables
		env := os.Getenv(strings.ToUpper(tag))

		// Log field information
		log.WithFields(log.Fields{
			"name":     value.Type().Field(i).Name,
			"kind":     kind,
			"tag":      tag,
			"flagName": flagName,
			"env":      env,
		}).Info("Parsing field")

		// Handle field by it's "kind"
		switch kind {
		case reflect.Bool:
			c.handleBoolField(value, i, flagName, env)
		case reflect.Int:
			c.handleIntField(value, i, flagName, env)
		case reflect.String:
			c.handleStringField(value, i, flagName, env)
		case reflect.Struct:
			// Recurse
			c.handleFields(value.Field(i).Addr())
		default:
			log.WithField("type", kind).Warn("Unable to handle field with the specified kind")
		}
	}
}

// handleBoolField handles fields with a "kind" of bool
// Sets a field's value as well as a flag (when allowed)
func (c *Config) handleBoolField(value reflect.Value, i int, flagName string, env string) {
	// Store field
	field := value.Field(i)

	// Handle non-empty enviroment variable
	if env != "" {
		parsed, _ := strconv.ParseBool(env)
		field.SetBool(parsed)
	}

	// If allowed, set a flag
	// NOTE: Checks PkgPath for empty value, meaning the field is exported
	// and thus reflect's Interface method can return it's value
	// See https://golang.org/pkg/reflect/#StructField for more information
	if flag.Lookup(flagName) == nil && value.Type().Field(i).PkgPath == "" {
		ptr := field.Addr().Interface().(*bool)
		flag.BoolVar(ptr, flagName, field.Bool(), "")
	}
}

// handleIntField handles fields with a "kind" of int
// Sets a field's value as well as a flag (when allowed)
func (c *Config) handleIntField(value reflect.Value, i int, flagName string, env string) {
	// Store field
	field := value.Field(i)

	// Handle non-empty enviroment variable
	if env != "" {
		parsed, _ := strconv.ParseInt(env, 10, 0)
		field.SetInt(int64(parsed))
	}

	// If allowed, set a flag
	// NOTE: Checks PkgPath for empty value, meaning the field is exported
	// and thus reflect's Interface method can return it's value
	// See https://golang.org/pkg/reflect/#StructField for more information
	if flag.Lookup(flagName) == nil && value.Type().Field(i).PkgPath == "" {
		ptr := field.Addr().Interface().(*int)
		flag.IntVar(ptr, flagName, int(field.Int()), "")
	}
}

// handleStringField handles fields with a "kind" of string
// Sets a field's value as well as a flag (when allowed)
func (c *Config) handleStringField(value reflect.Value, i int, flagName string, env string) {
	// Store field
	field := value.Field(i)

	// Handle non-empty enviroment variable
	if env != "" {
		field.SetString(env)
	}

	// If allowed, set a flag
	// NOTE: Checks PkgPath for empty value, meaning the field is exported
	// and thus reflect's Interface method can return it's value
	// See https://golang.org/pkg/reflect/#StructField for more information
	if flag.Lookup(flagName) == nil && value.Type().Field(i).PkgPath == "" {
		ptr := field.Addr().Interface().(*string)
		flag.StringVar(ptr, flagName, field.String(), "")
	}
}

// formFlagName converts a field's tag corresponding to an enviroment variable
// into a string to use as a flag's name. Will strip the application prefix
// and replace underscores with hypens. Will also return the name in lowercase.
// Ex: AEON_FOO_BAR_BAZ => foo-bar-baz
func (c *Config) formFlagName(temp string) string {
	// Form flag name
	name := strings.TrimPrefix(strings.ToUpper(temp), ENV_PREFIX)
	name = strings.Replace(name, "_", "-", -1)

	return strings.ToLower(name)
}
