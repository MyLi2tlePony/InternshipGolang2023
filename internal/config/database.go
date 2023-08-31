package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type databaseConfig struct {
	Prefix       string
	DatabaseName string
	Host         string
	Port         string
	UserName     string
	Password     string
}

func NewDatabaseConfig(configPath string) (*databaseConfig, error) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &databaseConfig{
		Prefix:       viper.GetString("database.Prefix"),
		DatabaseName: viper.GetString("database.DatabaseName"),
		Host:         viper.GetString("database.Host"),
		Port:         viper.GetString("database.Port"),
		UserName:     viper.GetString("database.UserName"),
		Password:     viper.GetString("database.Password"),
	}, nil
}

func (d *databaseConfig) GetPrefix() string {
	return d.Prefix
}

func (d *databaseConfig) GetDatabaseName() string {
	return d.DatabaseName
}

func (d *databaseConfig) GetHost() string {
	return d.Host
}

func (d *databaseConfig) GetPort() string {
	return d.Port
}

func (d *databaseConfig) GetUserName() string {
	return d.UserName
}

func (d *databaseConfig) GetPassword() string {
	return d.Password
}

func (d *databaseConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		d.Host, d.UserName, d.Password, d.DatabaseName, d.Port)
}
