# Dead Letter Exchange (DLX) in RabbitMQ

![Dead Letter Exchange](https://www.cloudamqp.com/img/blog/dead-letter-exchange-1.png)

A Dead Letter Exchange (DLX) is a special exchange that RabbitMQ uses to route messages that were not successfully processed by their original queue. Messages can become "dead-lettered" in several scenarios:

1. The message is rejected (basic.reject or basic.nack) with `requeue=false`.
2. The message expires due to Time-To-Live (TTL) constraints (either at the queue or per-message level).
3. The queue has reached its length limit, causing older messages to be dropped.

When any of these situations occur, the broker (RabbitMQ) will re-publish the "dead" message to a specified Dead Letter Exchange instead of discarding it. This makes it possible to handle these dead-lettered messages in a controlled manner (e.g., logging, error handling, alerting, or re-processing).

## How DLX works in RabbitMQ
1. Declare a Dead Letter Exchange.  
You need a dedicated exchange that will receive dead-lettered messages. Often, this is a topic or a direct exchange (depending on your routing needs).

2. Declare a Dead Letter Queue.    
Create a queue for storing the messages that are routed to the DLX. You bind this queue to the DLX with an appropriate routing key.

3. Declare your “main” queue with DLX arguments.    
When you declare the main queue (the one from which you are consuming normally), you add two (or one) important arguments:
    - `x-dead-letter-exchange`: the name of the DLX to which RabbitMQ will send the dead-lettered messages.
    - (Optionally) `x-dead-letter-routing-key`: the routing key to use when the message is re-published to the DLX. If you don't provide this, RabbitMQ will use the original routing key of the message.

4. Cause messages to be dead-lettered.  
If a consumer:
	- Rejects a message with `requeue=false` (via `basic.reject` or `basic.nack`),
	- The message expires (TTL), or
	- The queue is at capacity (if configured),

	then RabbitMQ will automatically move that message to the DLX.

5. Consumers on the DLX queue process the dead-lettered messages.   
You can have a separate consumer or processing logic to handle, inspect, or possibly re-process messages on the dead-letter queue.

Below is a simplified, self-contained example in Go demonstrating how to set up a dead-letter exchange and queue, along with a main queue that uses this DLX. It also shows publishing a message and then rejecting it so that it ends up in the dead-letter queue.

```go
package main
import (
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1. Connect to RabbitMQ
	conn, err := amqp091-go.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	// 2. Declare the Dead Letter Exchange
	deadLetterExchangeName := "example.dlx"
	err = channel.ExchangeDeclare(
		deadLetterExchangeName, // name
		"direct",               // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare the dead letter exchange")

	// 3. Declare a Dead Letter Queue
	deadLetterQueueName := "example.dlq"
	_, err = channel.QueueDeclare(
		deadLetterQueueName, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare the dead letter queue")

	// 4. Bind the Dead Letter Queue to the Dead Letter Exchange
	dlxRoutingKey := "dlx-key"
	err = channel.QueueBind(
		deadLetterQueueName,    // queue name
		dlxRoutingKey,          // routing key
		deadLetterExchangeName, // exchange
		false,                  // noWait
		nil,                    // args
	)
	failOnError(err, "Failed to bind the dead letter queue to the DLX")

	// 5. Declare the main exchange (optional)
	mainExchangeName := "example.exchange"
	err = channel.ExchangeDeclare(
		mainExchangeName, // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // noWait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare the main exchange")

	// 6. Declare the main queue with DLX arguments
	mainQueueName := "example.main.queue"
	args := amqp091-go.Table{
		"x-dead-letter-exchange":    deadLetterExchangeName,
		"x-dead-letter-routing-key": dlxRoutingKey,
		// "x-message-ttl": int32(5000), // optional TTL (5s) for demonstration
	}

	_, err = channel.QueueDeclare(
		mainQueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // noWait
		args,          // arguments
	)
	failOnError(err, "Failed to declare the main queue")

	// 7. Bind the main queue to the main exchange
	mainRoutingKey := "main-key"
	err = channel.QueueBind(
		mainQueueName,     // queue name
		mainRoutingKey,    // routing key
		mainExchangeName,  // exchange
		false,             // noWait
		nil,
	)
	failOnError(err, "Failed to bind the main queue")

	// 8. Publish a message to the main exchange
	messageBody := "Hello, this message might be dead-lettered"
	err = channel.PublishWithContext(
		// Using PublishWithContext allows context usage. You can also use Publish if you prefer.
		// If you don't want to use a context, use channel.Publish(...) directly.
		nil,               // context (nil for simplicity, or context.Background())
		mainExchangeName,  // exchange
		mainRoutingKey,    // routing key
		false,             // mandatory
		false,             // immediate
		amqp091-go.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		},
	)
	failOnError(err, "Failed to publish message to main queue")
	log.Printf("Published message to main queue: %s", messageBody)

	// 9. Consume from the main queue and force a reject to trigger DLX
	msgs, err := channel.Consume(
		mainQueueName, // queue
		"",            // consumer
		false,         // autoAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer on main queue")

	go func() {
		for d := range msgs {
			log.Printf("Received a message from main queue: %s", d.Body)
			// Reject (or Nack) the message so it goes to the DLQ
			if err := d.Nack(false, false); err != nil {
				log.Printf("Failed to Nack message: %v", err)
			} else {
				log.Println("Rejected the message; it should go to the DLQ.")
			}
		}
	}()

	// Give some time for the consumer to process and reject
	time.Sleep(2 * time.Second)

	// 10. Consume from the DLQ to verify the dead-lettered message
	dlqMsgs, err := channel.Consume(
		deadLetterQueueName, // queue
		"",                  // consumer
		true,                // autoAck (true for quick demo)
		false,               // exclusive
		false,               // noLocal
		false,               // noWait
		nil,                 // args
	)
	failOnError(err, "Failed to register a consumer on DLQ")

	go func() {
		for dlqMsg := range dlqMsgs {
			log.Printf("Received a message from DLQ: %s", dlqMsg.Body)
		}
	}()

	// Keep the program alive to observe the logs
	select {}
}
```

### Explanation
1. Dead Letter Exchange (DLX) & Queue   
We declare an exchange (`example.dlx`) that will receive any dead-lettered messages.
We declare a queue (`example.dlq`) bound to that exchange (with `dlx-key`).

2. Main Queue with DLX arguments    
- We declare a queue (`example.main.queue`) and set the arguments:
    - `x-dead-letter-exchange` to point to `example.dlx`
    - `x-dead-letter-routing-key` to `dlx-key` (optional; if omitted, original routing key is used).

3. Publish Message  
A message is published to `example.exchange` with routing key `main-key`.

4. Consume & Force Rejection   
We consume from `example.main.queue` and explicitly call `Nack` with `requeue=false`.
This will trigger RabbitMQ to send that message to the DLX.

5. Consume DLX Queue   
We consume from `example.dlq` to see that the message indeed arrived there after being dead-lettered.