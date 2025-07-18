package rabbitmq

import (
	"log"
	"order-service/package/config"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	url     string
}

func NewRabbitMQClient(cfg *config.Config) (*RabbitMQClient, error) {
	client := &RabbitMQClient{url: cfg.RabbitMQURL}
	if err := client.connect(); err != nil {
		return nil, err
	}
	return client, nil
}

func (r *RabbitMQClient) connect() error {
	var err error
	for i := 0; i < 5; i++ {
		r.conn, err = amqp.Dial(r.url)
		if err == nil {
			break
		}
		log.Printf("Retrying RabbitMQ connection... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return err
	}

	r.channel, err = r.conn.Channel()
	return err
}

func (r *RabbitMQClient) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
