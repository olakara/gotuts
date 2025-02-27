package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type JobCard struct {
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}

func main() {
	// Initialize RabbitMQ connection
	rabbitmq, err := NewRabbitMQ("amqp://guest:guest@localhost:5672/", "maint", true, false, false, false)
	FailOnError(err, "Failed to initialize RabbitMQ")
	defer rabbitmq.Close()

	// Set quality of service
	err = rabbitmq.SetQos(1, 0, false)
	FailOnError(err, "Failed to set QoS")

	// Start consuming messages
	msgs, err := rabbitmq.ConsumeMessages(false, false, false, false)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("starting job")

			var job JobCard
			if err := json.Unmarshal(d.Body, &job); err != nil {
				log.Printf("Failed to unmarshal job card: %v", err)
				continue
			}
			fmt.Println("Working on job:", job.Message, "with time:", job.Priority)
			// Simulate job processing
			time.Sleep(time.Duration(job.Priority) * time.Second)
			fmt.Println("completed job")
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
