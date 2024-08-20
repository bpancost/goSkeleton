package config

import (
	"encoding/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"time"
)

type Configuration struct {
	config *koanf.Koanf
}

// NewConfig Creates a new configuration loader using the Viper loader.
// The configuration is loaded from a file named "config" in yaml format in the cmd folder matching the given project name.
// Environment variables are loaded with precedence and must be prefixed by the project name in snake case (all caps).
func NewConfig(projectName string) (Config, error) {
	k := koanf.New(".")
	err := k.Load(file.Provider("config.yaml"), yaml.Parser()) // Running locally
	errCount := 0
	if err != nil {
		errCount++
	}
	err = k.Load(file.Provider("cmd/"+projectName+"/config.yaml"), yaml.Parser()) // Running locally
	if err != nil {
		errCount++
	}
	err = k.Load(file.Provider("bin/"+projectName+"/config.yaml"), yaml.Parser()) // Running the dev container image
	if err != nil {
		errCount++
	}

	err = k.Load(env.Provider(toSnakeCase(projectName)+"_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, toSnakeCase(projectName)+"_")), "_", ".", -1)
	}), nil)
	if err != nil {
		errCount++
	}
	var outError error
	if errCount == 4 {
		outError = errors.New("No configuration was loaded")
	}

	return &Configuration{
		config: k,
	}, outError
}

func (c *Configuration) GetBoolean(key string) bool {
	return c.config.Bool(key)
}

func (c *Configuration) GetString(key string) string {
	return c.config.String(key)
}

func (c *Configuration) GetInt(key string) int {
	return c.config.Int(key)
}

func (c *Configuration) GetInt32(key string) int32 {
	return int32(c.config.Int(key))
}

func (c *Configuration) GetInt64(key string) int64 {
	return c.config.Int64(key)
}

func (c *Configuration) GetFloat64(key string) float64 {
	return c.config.Float64(key)
}

func (c *Configuration) GetTime(key string) time.Time {
	return c.config.Time(key, time.RFC3339)
}

func (c *Configuration) GetDuration(key string) time.Duration {
	return c.config.Duration(key)
}

func (c *Configuration) GetStringMapString(key string) map[string]string {
	return c.config.StringMap(key)
}

func (c *Configuration) Get(key string) interface{} {
	return c.config.Get(key)
}

func (c *Configuration) GetListOfMaps(key string) ([]map[string]string, error) {
	finalValues := make([]map[string]string, 0)
	rawValue := c.config.Get(key)
	if listValues, ok := rawValue.([]interface{}); ok {
		finalValues = make([]map[string]string, len(listValues))
		for index, objectRaw := range listValues {
			if objectMap, ok := objectRaw.(map[interface{}]interface{}); ok {
				stringMap := make(map[string]string)
				for key, value := range objectMap {
					keyString, okKey := key.(string)
					keyValue, okValue := value.(string)
					if okKey && okValue {
						stringMap[keyString] = keyValue
					}
				}
				finalValues[index] = stringMap
			}
		}
	} else if stringValue, ok := rawValue.(string); ok {
		err := json.Unmarshal([]byte(stringValue), &finalValues)
		return finalValues, err
	}
	return finalValues, nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
