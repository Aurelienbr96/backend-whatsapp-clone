package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Twilio   TwilioConfig
	Postgres PostgresConfig
	RabbitMq RabbitMqConfig
}

type TwilioConfig struct {
	TwilioAccountSid string
	TwilioAuthToken  string
	VerifyServiceSid string
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Sslmode  string
}

type RabbitMqConfig struct {
	Servers string
}

func LoadConfig() (*viper.Viper, error) {
	config := viper.New()

	config.SetConfigName("config-local")
	config.AddConfigPath("config")
	config.AddConfigPath("../config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return config, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
