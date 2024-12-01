package queue

import (
	"context"
	"fmt"
	"time"

	e "example.com/boiletplate/internal/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

//go:generate mockgen -source=publisher.go -destination=mock/publisher.go

type IPublisher interface {
	PushMessage(body []byte) error
}

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

func (p Publisher) PushMessage(body []byte) error {
	fmt.Println(
		" [x] Sent %s",
		body,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	err = p.ch.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}
