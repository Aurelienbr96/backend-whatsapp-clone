package rabbitmq

import (
	"fmt"

	"example.com/boiletplate/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMq(config *config.Config) (*amqp.Connection, error) {

	connString := fmt.Sprintf("amqp://guest:guest@%s", config.RabbitMq.Servers)
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
