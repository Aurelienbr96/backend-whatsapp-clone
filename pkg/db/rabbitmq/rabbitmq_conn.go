package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"example.com/boiletplate/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMq(config *config.Config) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	connString := fmt.Sprintf("amqp://guest:guest@%s", config.RabbitMq.Servers)

	for i := 0; i < 10; i++ { // Retry up to 10 times
		conn, err = amqp.Dial(connString)
		if err == nil {
			return conn, nil
		}
		log.Printf("RabbitMQ is not ready yet, retrying in 5 seconds... (%d/10)", i+1)
		time.Sleep(5 * time.Second)
	}

	return nil, err
}
