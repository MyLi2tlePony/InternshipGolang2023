package config

import (
	"net"

	"github.com/spf13/viper"
)

type serverConfig struct {
	Host string
	Port string
}

func NewServerConfig(configPath string) (*serverConfig, error) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &serverConfig{
		Host: viper.GetString("server.Host"),
		Port: viper.GetString("server.Port"),
	}, nil
}

func (s *serverConfig) GetPort() string {
	return s.Port
}

func (s *serverConfig) GetHost() string {
	return s.Host
}

func (s *serverConfig) GetHostPort() string {
	return net.JoinHostPort(s.Host, s.Port)
}
