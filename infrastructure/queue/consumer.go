package queue

import (
	"encoding/json"
	"log"

	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	e "example.com/boiletplate/internal/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	ch         *amqp.Channel
	otpHandler otphandler.OTPHandler
}

func NewConsumer(conn *amqp.Connection, otpHandler otphandler.OTPHandler) *Consumer {
	ch, err := conn.Channel()
	e.FailOnError(err, "Failed to open a channel for publisher")
	return &Consumer{ch: ch, otpHandler: otpHandler}
}

func (c *Consumer) Subscribe() {
	err := c.ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	e.FailOnError(err, "Failed to declare an exchange")

	q, err := c.ch.QueueDeclare(
		"t",   // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	e.FailOnError(err, "Failed to declare a queue")

	err = c.ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	e.FailOnError(err, "Failed to bind a queue")

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	e.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msg := CreatedUserSuccess{}
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				e.FailOnError(err, "Could not process the message")
			}

			if msg.Type == "created_user" && msg.Payload.PhoneNumber != "" {
				c.otpHandler.SendOTP(msg.Payload.PhoneNumber)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
