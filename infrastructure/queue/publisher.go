package queue

import (
	"context"
	"github.com/stretchr/testify/mock"
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

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PushMessage(body []byte) error {
	args := m.Called(body)
	return args.Error(0)
}

func (p Publisher) PushMessage(body []byte) error {
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
