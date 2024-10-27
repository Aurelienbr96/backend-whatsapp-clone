package queue

import (
	"context"
	"log"
	"time"

	e "example.com/boiletplate/internal/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) *Publisher {
	ch, err := conn.Channel()
	e.FailOnError(err, "Failed to open a channel for publisher")
	return &Publisher{
		ch: ch,
	}
}

func (p Publisher) PushMessage(body []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p.ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	err := p.ch.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		e.FailOnError(err, "Failed to push message to the queue")
	}
	log.Printf(" [x] Sent %s", body)
}
