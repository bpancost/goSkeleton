package config

import (
	"github.com/spf13/viper"
	"regexp"
	"strings"
	"time"
)

type ViperConfig struct {
	config *viper.Viper
}

func NewViperConfig(projectName string) (Config, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./cmd/" + projectName)
	config.AddConfigPath(".")
	config.SetEnvPrefix(toSnakeCase(projectName))
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &ViperConfig{
		config: config,
	}, nil
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

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
