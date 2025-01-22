package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// failOnError is a helper function to check for errors and log them if they occur.
func failOnError(err error, msg string) {
	if err != nil {
		// If an error exists, log it with the given message
		log.Panicf("%s: %s", msg, err)
	}
}

// consumeMessages listens for and consumes messages from the specified queue
func consumeMessages(ch *amqp.Channel, queueName string) {
	// When you call ch.Consume(), it sets up an internal background consumer in RabbitMQ. This internal consumer listens to the specified queue.
	// As messages arrive, they are pushed to the consumer (the goroutine that is listening for messages), and your for loop processes them.
	msgs, err := ch.Consume(
		queueName, // Queue name to consume from
		"",        // Consumer tag: empty string means no specific consumer tag
		false,     // Auto-ack: Automatically acknowledge messages once received (In the context of auto-acknowledgment (true in ch.Consume), you don’t need to use d.Ack(), as RabbitMQ will automatically handle message acknowledgment for you when the message is delivered to the consumer.)
		false,     // Exclusive: This consumer can share the queue with others (false means it can share)
		false,     // No-local: Disallow messages from being sent to the same connection (false means allowed)
		false,     // No-wait: Do not wait for a response from the server (false means wait)
		nil,       // Arguments: Additional arguments can be provided (nil means no arguments)
	)

	// Check if there was an error registering the consumer
	failOnError(err, "Failed to register a consumer")

	// Since it will push us messages asynchronously, we will read the messages from a channel (returned by amqp::Consume) in a goroutine.
	go func() {
		// Loop over the received messages
		for d := range msgs {
			// Log the received message
			log.Printf("Received a message: %s", d.Body)
			// Either Delivery.Ack, Delivery.Reject or Delivery.Nack must be called for every delivery that is not automatically acknowledged.
			// When multiple is true, this delivery and all prior unacknowledged deliveries on the same channel will be acknowledged. This is useful for batch processing of deliveries.
			d.Ack(false)
			/*
				It's a common mistake to miss the ack. It's an easy error, but the consequences are serious. Messages will be redelivered when your client quits (which may look like random redelivery), but RabbitMQ will eat more and more memory as it won't be able to release any unacked messages.
			*/

			/*
				d.Nack(multiple bool, requeue bool):
					When multiple is true, nack messages up to and including delivered messages up until the delivery tag delivered on the same channel.

					When requeue is true, request the server to deliver this message to a different consumer. If it is not possible or requeue is false, the message will be dropped or delivered to a server configured dead-letter queue.

					This method must not be used to select or requeue messages the client wishes not to handle, rather it is to inform the server that the client is incapable of handling this message at this time.
			*/
			/*
				d.Reject(requeue bool):
					When requeue is true, queue this message to be delivered to a consumer on a different channel. When requeue is false or the server is unable to queue this message, it will be dropped.

					If you are batch processing deliveries, and your server supports it, prefer Delivery.Nack.
			*/

			// In both the Delivery.Nack and Delivery.Reject methods, when requeuing a message, there are two approaches we can take if the message is received again. One option is to add a max_try header to the message, allowing us to track how many times it has been processed. The other option is to set a TTL (Time-To-Live) for the messages in the queue when declaring the queue, which would limit the time the message remains in the queue before being discarded.
		}
	}()

	// Log that the program is waiting for messages
	log.Printf(" [*] waiting for messages. To exit press CTRL+C")
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		failOnError(conn.Close(), "Failed to Close the connection")
	}()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() {
		failOnError(ch.Close(), "Failed to Close the channel")
	}()

	// Note that we declare the queue here, as well. Because we might start the consumer before the publisher, we want to make sure the queue exists before we try to consume messages from it.
	queueName := "test_queue"
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		amqp.Table{
			// "x-queue-type": "classic", // classic by default. another type is quorom
			// "x-max-length":  10,       // Maximum length on a queue
			// "x-message-ttl": 10_000,   // message TTL
		},
	)

	// Check if there was an error declaring the queue
	failOnError(err, "Failed to declare a queue")

	// Start consuming messages from the queue in a separate goroutine
	consumeMessages(ch, q.Name)

	// This will block the main function indefinitely, allowing the consumer to keep working
	select {}
}
