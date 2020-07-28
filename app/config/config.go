package config

import "time"

// A wrapper for config loaders
type Config interface {
	// Get a boolean value from the config using the given key
	GetBoolean(key string) bool
	// Get a string value from the config using the given key
	GetString(key string) string
	// Get an int value from the config using the given key
	GetInt(key string) int
	// Get an int32 value from the config using the given key
	GetInt32(key string) int32
	// Get an int64 value from the config using the given key
	GetInt64(key string) int64
	// Get a float64 value from the config using the given key
	GetFloat64(key string) float64
	// Get a time value from the config using the given key
	GetTime(key string) time.Time
	// Get a duration value from the config using the given key
	GetDuration(key string) time.Duration
	// Get a map of strings keys to string values from the config using the given key
	GetStringMapString(key string) map[string]string
	// Get a value without type information from the config using the given key
	Get(key string) interface{}
}
