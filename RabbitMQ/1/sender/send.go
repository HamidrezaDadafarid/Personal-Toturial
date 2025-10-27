package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go" // Import RabbitMQ library (amqp091-go is a client for AMQP 0-9-1)
)

// failOnError is a helper function to check for errors and log them if they occur.
func failOnError(err error, msg string) {
	if err != nil {
		// If an error exists, log it with the given message
		log.Panicf("%s: %s", msg, err)
	}
}

// publishMessage publishes a message to the specified queue
func publishMessage(ch *amqp.Channel, queueName string, body string) {
	// Publishing the message to the RabbitMQ queue
	err := ch.Publish(
		"",        // Exchange: empty string means direct communication with the queue (a special Direct Exchange with an empty name)
		queueName, // Routing key: This is the queue name to which the message will be sent
		false,     // Mandatory: Message will be returned if it's not routed to any queue
		false,     // Immediate: If the message cannot be routed immediately, it will be discarded
		amqp.Publishing{
			ContentType: "text/plain", // Message type, typically "text/plain" for simple text messages
			Body:        []byte(body), // The message body content (in this case, the "Hello RabbitMQ!" string)
			// Expiration:  "5000",    // TTL per message in miliseconds. When both a per-queue and a per-message TTL are specified, the lower value between the two will be chosen.
		},
	)

	// Check if there was an error publishing the message
	failOnError(err, "Failed to publish a message")
	// Log the message sent
	log.Printf(" [x] sent %s", body)
}

func main() {
	// Establish a connection to the RabbitMQ server at localhost on the default port 5672
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // AMQP connection string (default credentials)
	failOnError(err, "Failed to connect to RabbitMQ")            // Check for connection error
	defer func() {
		failOnError(conn.Close(), "Failed to Close the connection")
	}()

	// Create a new channel over the connection
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel") // Check for channel creation error
	defer func() {
		failOnError(ch.Close(), "Failed to Close the channel")
	}()

	// Declare a queue to use. This ensures the queue exists before publishing/consuming messages
	// Declaring a queue is idempotent - it will only be created if it doesn't exist already.
	queueName := "test_queue"
	// We cannot change the features and properties of the queue when the queue is up and declared. we have to delete the queue and recreate the queue if we want to change its properties.
	q, err := ch.QueueDeclare(
		queueName, // Queue name
		false,     // Durable: If true, queue survives server restarts, false means it doesn't survive
		false,     // Delete when unused: If true, the queue is deleted when no consumers are attached
		false,     // Exclusive: The queue used by only one connection and the queue will be deleted when that connection closes
		false,     // No-wait: If true, server will not send a response after declaring the queue
		// Arguments: Additional queue-specific configurations (empty for now)
		amqp.Table{
			// You can define additional properties here for the queue
			// "x-queue-type": "classic", // classic by default. another type is quorom
			// "x-max-length":  10,       // Maximum length on a queue
			// "x-message-ttl": 10_000,   // message TTL
		},
	)

	// Log the name of the declared queue
	fmt.Println(q.Name)
	// Check if there was an error declaring the queue
	failOnError(err, "Failed to declare a queue")

	// Message body to publish to the queue
	body := "Hello RabbitMQ!"

	// Publish a message to the queue
	publishMessage(ch, queueName, body)

	// Block indefinitely to keep the main function running (otherwise the program will exit immediately)
	select {}
}
