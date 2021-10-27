package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Service - defines the interface with methods that our
// struct must contains
type Service interface {
	Connect() error
	Publish(message string) error
	Consume() error
}

// RabbitMQ - defines a struct which contains things like
// pointers to amqp connection and methods
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Connect - establishes a connection to our RabbitMQ instance
// and declares the queue we are going to be using
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error

	r.Conn, err = amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to RabbitMQ")

	// We need to open a channel over our AMQP connection
	// This will allow us to declare queue and subsequently
	// consume/publish messages
	r.Channel, err = r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Here we declare out new queue that we want to publish to
	// and consume from:
	_, err = r.Channel.QueueDeclare(
		"TestQueue", // Queue name
		false,       // durable
		false,       // Delete when not used
		false,       // exclusive
		false,       // no wait
		nil,         // additional args
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Publish - publishes a message to queue
func (r *RabbitMQ) Publish(message string) error {
	// attempt to publish a message to queue!
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Successfully published message to queue")
	return nil
}

// Consume - consumes messages from our test queue
func (r *RabbitMQ) Consume() error {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received Message: %s\n", msg.Body)
	}

	return nil
}

// NewRabbitMQService - returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
