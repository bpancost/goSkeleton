package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"regexp"
	"strings"
	"time"
)

type ViperConfig struct {
	config *viper.Viper
}

// NewViperConfig Creates a new configuration loader using the Viper loader.
// The configuration is loaded from a file named "config" in yaml format in the cmd folder matching the given project name.
// Environment variables are loaded with precedence and must be prefixed by the project name in snake case (all caps).
func NewViperConfig(projectName string) (Config, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./cmd/" + projectName) // running locally
	config.AddConfigPath(".")                    // running locally
	config.AddConfigPath("./bin/")               // running the dev container image
	config.SetEnvPrefix(toSnakeCase(projectName))
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()
	err := config.ReadInConfig()
	return &ViperConfig{
		config: config,
	}, err
}

func (c *ViperConfig) GetBoolean(key string) bool {
	return c.config.GetBool(key)
}

func (c *ViperConfig) GetString(key string) string {
	return c.config.GetString(key)
}

func (c *ViperConfig) GetInt(key string) int {
	return c.config.GetInt(key)
}

func (c *ViperConfig) GetInt32(key string) int32 {
	return c.config.GetInt32(key)
}

func (c *ViperConfig) GetInt64(key string) int64 {
	return c.config.GetInt64(key)
}

func (c *ViperConfig) GetFloat64(key string) float64 {
	return c.config.GetFloat64(key)
}

func (c *ViperConfig) GetTime(key string) time.Time {
	return c.config.GetTime(key)
}

func (c *ViperConfig) GetDuration(key string) time.Duration {
	return c.config.GetDuration(key)
}

func (c *ViperConfig) GetStringMapString(key string) map[string]string {
	return c.config.GetStringMapString(key)
}

func (c *ViperConfig) Get(key string) interface{} {
	return c.config.Get(key)
}

func (c *ViperConfig) GetListOfMaps(key string) ([]map[string]string, error) {
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
