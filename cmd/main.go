package main

import (
	"example.com/boiletplate/config"
	"example.com/boiletplate/pkg/db/postgres"
	"example.com/boiletplate/pkg/db/rabbitmq"
	twilio_client "example.com/boiletplate/pkg/otp_provider/twilio"

	"example.com/boiletplate/internal/errors"
	"example.com/boiletplate/internal/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	viper, err := config.LoadConfig()
	errors.FailOnError(err, "Could not load viper config")

	config, err := config.ParseConfig(viper)
	errors.FailOnError(err, "Could not parse viper config")

	entClient, err := postgres.NewPsqlDB(config)
	errors.FailOnError(err, "Could not set connect to postgresql")
	defer entClient.Close()

	rabbitMqClient, err := rabbitmq.NewRabbitMq(config)
	errors.FailOnError(err, "Could not set connect to rabbitmq")
	defer rabbitMqClient.Close()

	twilioClient := twilio_client.NewTwilioClient(config)
	server := server.NewServer(entClient, rabbitMqClient, twilioClient)
	server.Bootstrap()
}
