package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// JobCard represents a maintenance job with a message and priority
type JobCard struct {
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Check command-line arguments
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <message> <priority>", os.Args[0])
	}

	message := os.Args[1]
	priority, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to parse priority as integer")

	// Create a new JobCard
	job := JobCard{
		Message:  message,
		Priority: priority,
	}

	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"maint", // queue name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Convert JobCard to JSON
	body, err := json.Marshal(job)
	failOnError(err, "Failed to encode JobCard as JSON")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish the JobCard to the queue
	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	failOnError(err, "Failed to publish message")

	fmt.Printf("Sent JobCard with message '%s' and priority %d to queue 'maint'\n", message, priority)
}
