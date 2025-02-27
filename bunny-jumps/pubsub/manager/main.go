package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strconv"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	// check command line arguments
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <title> <duration>", os.Args[0])
	}

	title := os.Args[1]
	duration, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to convert duration to integer")

	// create a new job card
	jobCard := NewJobCard(title, duration)

	// connect to Rabbit MQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declare a queue
	q, err := ch.QueueDeclare(
		"jobs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// convert jobcard to json
	message, err := json.Marshal(jobCard)
	failOnError(err, "Failed to convert jobcard to json")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// publish the job card to the queue
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         message,
		})
	failOnError(err, "Failed to publish a message")
	fmt.Println(" [x] Sent ", jobCard.Title)
}
