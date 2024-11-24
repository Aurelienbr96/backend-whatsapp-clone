package main

import (
	"fmt"
	"log"

	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"

	"example.com/boiletplate/config"
	"example.com/boiletplate/pkg/db/postgres"
	"example.com/boiletplate/pkg/db/rabbitmq"
	twilioclient "example.com/boiletplate/pkg/otp_provider/twilio"

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
	log.Printf("ent client: %v", entClient)
	errors.FailOnError(err, "Could not set connect to postgresql")
	defer entClient.Close()

	fmt.Println("Connected to postgres")

	rabbitMqClient, err := rabbitmq.NewRabbitMq(config)
	errors.FailOnError(err, "Could not set connect to rabbitmq")
	defer rabbitMqClient.Close()

	twilioClient := twilioclient.NewTwilioClient(config)
	twilioAdapter := otphandler.NewTwilioAdapter(twilioClient.Twilio, twilioClient.VerifyServiceSid)

	consumer := queue.NewConsumer(rabbitMqClient, twilioAdapter)
	go consumer.Subscribe()

	publisher := queue.NewPublisher(rabbitMqClient)

	s := server.NewServer(entClient, twilioAdapter, publisher)
	s.Bootstrap()
}
