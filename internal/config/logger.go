package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type loggerConfig struct {
	Level string
}

func NewLoggerConfig(configPath string) (*loggerConfig, error) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &loggerConfig{
		Level: viper.GetString("logger.level"),
	}, nil
}

func (l loggerConfig) GetLevel() log.Lvl {
	switch l.Level {
	case "info":
		return log.INFO
	case "debug":
		return log.DEBUG
	case "error":
		return log.ERROR
	case "warn":
		return log.WARN
	}

	return log.OFF
}
