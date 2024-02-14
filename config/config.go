package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig
	Nats   NatsConfig
}

type ServerConfig struct {
	AppVersion string `json:"appVersion"`
	Host       string `json:"host" validate:"required"`
	Port       string `json:"port" validate:"required"`
}

type NatsConfig struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Cluster string `json:"cluster"`
	Client  string `json:"client"`
	Topic   string `json:"topic"`
}

func LoadConfig() (*viper.Viper, error) {

	viperInstance := viper.New()

	viperInstance.AddConfigPath("./config")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yaml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
