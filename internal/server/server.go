package server

import (
	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	twilio_client "example.com/boiletplate/pkg/otp_provider/twilio"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

type Server struct {
	entClient      *ent.Client
	rabbitMqClient *amqp091.Connection
	twilioClient   *twilio_client.TwilioClient
	gin            *gin.Engine
}

func NewServer(entClient *ent.Client, rabbitMqClient *amqp091.Connection, twilioClient *twilio_client.TwilioClient) *Server {
	return &Server{gin: gin.Default(), entClient: entClient, rabbitMqClient: rabbitMqClient, twilioClient: twilioClient}
}

func (s *Server) Bootstrap() {

	twilioAdapter := otphandler.NewTwilioAdapter(s.twilioClient.Twilio, s.twilioClient.VerifyServiceSid)

	consumer := queue.NewConsumer(s.rabbitMqClient, twilioAdapter)
	go consumer.Subscribe()

	publisher := queue.NewPublisher(s.rabbitMqClient)

	NewHandlers(s, publisher)
	s.gin.Run()
}
