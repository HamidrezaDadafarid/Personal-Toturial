package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// failOnError is a helper function to log and panic if an error occurs.
// The msg parameter provides context about what operation failed.
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// publishMessage publishes a message to the specified exchange using a routing key.
func publishMessage(ch *amqp.Channel, exchangeName string, routingKey string, body string) {
	// Publish the message on the given channel to the specified exchange with a routing key.
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("[x] Sent %s", body)
}

// consumeMessages declares a consumer on the specified queue.
// It listens for messages in the background and logs them as they arrive.
func consumeMessages(ch *amqp.Channel, queueName string) {
	// ch.Consume returns a channel (msgs) that you can range over to read incoming messages.
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer name (empty string => auto-generated)
		true,      // auto-ack (messages are automatically considered acknowledged)
		false,     // exclusive (if true, only this consumer can access the queue)
		false,     // no-local (not supported by RabbitMQ, leave as false)
		false,     // no-wait (if true, do not wait for the server's response)
		nil,       // arguments
	)
	failOnError(err, "Failed to register a consumer")

	// We start a goroutine that will range over the msgs channel
	// and handle incoming messages. This allows the main goroutine to
	// continue doing other work (or in this case, block on select).
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
}

func main() {
	// 1. Establish a connection to RabbitMQ using the amqp.Dial function.
	//    The connection string includes protocol, username, password, host, port,
	//    and optionally a virtual host (the trailing '/').
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// Defer closing the connection until the main function returns.
	defer func() {
		err := conn.Close()
		if err != nil {
			failOnError(err, "Failed to close the connection")
		}
	}()

	// 2. Open a channel over the established connection.
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// Defer closing the channel until the main function returns.
	defer func() {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Failed to close the channel")
		}
	}()

	// 3. Declare an exchange. An exchange is responsible for receiving messages
	//    and routing them to queues based on the type of exchange and bindings.
	//    Here, we declare a "direct" exchange which routes messages based on an exact routing key match.
	exchangeName := "test_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		"direct",     // type of exchange (direct, fanout, topic, headers)
		true,         // durable (the exchange survives server restarts)
		false,        // auto-deleted (if true, the exchange is deleted when there are no bindings)
		false,        // internal (if true, the exchange cannot be directly published to by clients)
		false,        // no-wait (if true, declare without waiting for a server response)
		nil,          // arguments (optional custom arguments)
	)
	failOnError(err, "Failed to declare an exchange")

	// 4. Declare a queue where messages will be stored.
	//    If a queue with this name does not exist, RabbitMQ will create it.
	queueName := "test_queue"
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable (if true, the queue will survive server restarts)
		false,     // delete when unused (if true, it deletes itself when no consumers are connected)
		false,     // exclusive (if true, only this connection can use the queue)
		false,     // no-wait (if true, don't wait for a server response)
		nil,       // arguments (optional custom arguments, e.g. message TTL)
	)
	failOnError(err, "Failed to declare a queue")

	// 5. Bind the queue to the exchange with a specific routing key.
	//    This tells the exchange to route messages with the given routing key to the specified queue.
	routingKey := "test_key"
	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key to bind on
		exchangeName, // exchange name
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to bind the queue to the exchange")

	// 6. Publish a message to the exchange. The exchange will route it to
	//    any queues that are bound with a matching routing key.
	body := "Hello RabbitMQ via Exchange!"
	publishMessage(ch, exchangeName, routingKey, body)

	// 7. Set up a consumer to listen on the queue for messages.
	consumeMessages(ch, q.Name)

	// 8. Block forever (or until you manually stop the program) so the consumer
	//    goroutine can continue receiving messages.
	select {}
}
