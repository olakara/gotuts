package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ connection and channel structure
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewRabbitMQ initializes a new RabbitMQ connection and channel
func NewRabbitMQ(url string, queueName string, durable, autoDelete, exclusive, noWait bool) (*RabbitMQ, error) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Declare a queue
	q, err := ch.QueueDeclare(
		queueName,  // queue name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		nil,        // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

// SetQos sets the prefetch count for the channel
func (r *RabbitMQ) SetQos(prefetchCount, prefetchSize int, global bool) error {
	return r.channel.Qos(
		prefetchCount, // prefetch count
		prefetchSize,  // prefetch size
		global,        // global
	)
}

// ConsumeMessages starts consuming messages from the queue
func (r *RabbitMQ) ConsumeMessages(autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		autoAck,      // auto-ack
		exclusive,    // exclusive
		noLocal,      // no-local
		noWait,       // no-wait
		nil,          // args
	)
}

// Close closes the channel and connection
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// FailOnError is a utility function for error handling
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
