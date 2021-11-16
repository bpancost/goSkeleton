package config

import "time"

// Config A wrapper for config loaders
type Config interface {
	// GetBoolean Get a boolean value from the config using the given key
	GetBoolean(key string) bool
	// GetString Get a string value from the config using the given key
	GetString(key string) string
	// GetInt Get an int value from the config using the given key
	GetInt(key string) int
	// GetInt32 Get an int32 value from the config using the given key
	GetInt32(key string) int32
	// GetInt64 Get an int64 value from the config using the given key
	GetInt64(key string) int64
	// GetFloat64 Get a float64 value from the config using the given key
	GetFloat64(key string) float64
	// GetTime Get a time value from the config using the given key
	GetTime(key string) time.Time
	// GetDuration Get a duration value from the config using the given key
	GetDuration(key string) time.Duration
	// GetStringMapString Get a map of strings keys to string values from the config using the given key
	GetStringMapString(key string) map[string]string
	// Get a value without type information from the config using the given key
	Get(key string) interface{}
	// GetListOfMaps Gets a list of objects which each are translated to a map of string key to string value
	GetListOfMaps(key string) ([]map[string]string, error)
}
