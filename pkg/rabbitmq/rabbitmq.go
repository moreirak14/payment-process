package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	// Open a connection to RabbitMQ
	connection, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// Open a channel
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consume(channel *amqp.Channel, queueName string, output chan amqp.Delivery) error {
	// Consume from the queue
	messages, err := channel.Consume(
		queueName,  // queue
		"payments", // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		panic(err)
	}

	for message := range messages {
		output <- message
	}

	return nil
}

func Publish(ctx context.Context, channel *amqp.Channel, body, exName string) error {
	// Publish a message
	err := channel.PublishWithContext(
		ctx,
		exName, // exchange
		"",     // key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		panic(err)
	}

	return nil
}
