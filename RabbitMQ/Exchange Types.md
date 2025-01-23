# 1. Direct Exchange

![Direct Exchange](https://www.cloudamqp.com/img/blog/direct-exchange.svg)

A direct exchange routes messages to queues by matching the message’s routing key exactly against the queue binding’s routing key.

Use Case: When you need a 1:1 mapping between a specific routing key and a queue. For example, routing logs by severity where each severity (info, warning, error) is a unique routing key.

```go
package main

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // 1. Connect to RabbitMQ
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // 2. Create a channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // 3. Declare a direct exchange
    exchangeName := "direct_logs"
    err = ch.ExchangeDeclare(
        exchangeName, // name
        "direct",     // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    failOnError(err, "Failed to declare direct exchange")

    // 4. Declare a queue
    queueName := "direct_log_queue"
    q, err := ch.QueueDeclare(
        queueName, // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    failOnError(err, "Failed to declare queue")

    // 5. Bind queue to the exchange with a routing key
    //    (e.g., "info")
    routingKey := "info"
    err = ch.QueueBind(
        q.Name,
        routingKey,     // routing key
        exchangeName,   // exchange
        false,
        nil,
    )
    failOnError(err, "Failed to bind queue to direct exchange")

    // 6. Publish a message using the same routing key
    body := "Hello Direct Exchange!"
    err = ch.Publish(
        exchangeName,   // exchange
        routingKey,     // routing key
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    failOnError(err, "Failed to publish a message")

    log.Printf("[x] Sent %s with routing key '%s'", body, routingKey)

    // 7. Consume messages
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register consumer")

    go func() {
        for d := range msgs {
            log.Printf("Received: %s", d.Body)
        }
    }()

    log.Println("[*] Waiting for messages. To exit press CTRL+C")
    select {}
}
```
## How It Works
- Declare a direct exchange named direct_logs.
- Declare a queue named direct_log_queue.
- Bind the queue to the exchange with a specific routing key, e.g., "info".
- When publishing, you specify the same routing key in ch.Publish.
- Only the queue bound with that matching key receives the message.

# 2. Fanout Exchange

![Fanout Exchange](https://www.cloudamqp.com/img/blog/fanout-exchange.svg)

A fanout exchange routes messages to all bound queues, ignoring the routing key.

Use Case: When you want to broadcast a message to multiple consumers simultaneously. For example, sending out notifications or events that every service should receive.

```go
package main

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // 1. Connect to RabbitMQ
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // 2. Create a channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // 3. Declare a fanout exchange
    exchangeName := "fanout_logs"
    err = ch.ExchangeDeclare(
        exchangeName,
        "fanout", // type
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare fanout exchange")

    // 4. Declare a queue
    queueName := "fanout_log_queue"
    q, err := ch.QueueDeclare(
        queueName,
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare queue")

    // 5. Bind the queue to the fanout exchange.
    //    The routing key is ignored by a fanout exchange, but an empty string is usually given.
    err = ch.QueueBind(
        q.Name,
        "",           // routing key (ignored)
        exchangeName, // exchange
        false,
        nil,
    )
    failOnError(err, "Failed to bind queue to fanout exchange")

    // 6. Publish a message
    body := "Hello Fanout Exchange!"
    err = ch.Publish(
        exchangeName, // exchange
        "",           // routing key is often unused in a fanout
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    failOnError(err, "Failed to publish a message")

    log.Printf("[x] Sent %s", body)

    // 7. Consume messages
    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to register consumer")

    go func() {
        for d := range msgs {
            log.Printf("Received: %s", d.Body)
        }
    }()

    log.Println("[*] Waiting for messages. To exit press CTRL+C")
    select {}
}
```
## How It Works
- Any message sent to a fanout exchange is delivered to all queues bound to that exchange, regardless of a routing key.
- This is like a broadcast mechanism.

# 3. Topic Exchange

![Topic Exchange](https://www.cloudamqp.com/img/blog/topic-exchange.svg)

A topic exchange routes messages to one or many queues based on matching the routing key to a pattern. The routing key is usually a dot-separated string (e.g., "quick.orange.rabbit"). The binding keys (patterns) can include:

`*` (asterisk) matches exactly one word.

`#` (hash) matches zero or more words.

Use Case: Publish messages with multiple “categories,” “topics,” or “tags,” and selectively subscribe using patterns. For example, routing logs by severity and source: "kern.info", "auth.error", etc.

```go
package main

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // 1. Connect to RabbitMQ
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // 2. Create a channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // 3. Declare a topic exchange
    exchangeName := "topic_logs"
    err = ch.ExchangeDeclare(
        exchangeName,
        "topic", // type
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare topic exchange")

    // 4. Declare a queue
    //    This queue will receive messages that match our binding key pattern.
    queueName := "topic_log_queue"
    q, err := ch.QueueDeclare(
        queueName,
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare queue")

    // 5. Bind the queue with a pattern. 
    //    For example: 
    //    - "*.critical" matches any routing key with 'critical' as the second word.
    //    - "#.error" matches any routing key that ends with 'error', with zero or more words before it.
    bindPattern := "*.critical"
    err = ch.QueueBind(
        q.Name,
        bindPattern,
        exchangeName,
        false,
        nil,
    )
    failOnError(err, "Failed to bind queue to topic exchange")

    // 6. Publish a message with a routing key
    routingKey := "system.critical"
    body := "Hello Topic Exchange!"
    err = ch.Publish(
        exchangeName, 
        routingKey, // e.g., "system.critical"
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    failOnError(err, "Failed to publish a message")

    log.Printf("[x] Sent %s with routing key %s", body, routingKey)

    // 7. Consume messages
    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to register consumer")

    go func() {
        for d := range msgs {
            log.Printf("Received: %s", d.Body)
        }
    }()

    log.Println("[*] Waiting for messages. To exit press CTRL+C")
    select {}
}
```

## How It Works
- When binding a queue, you provide a pattern (e.g., "*.critical").
- Any message whose routing key matches the pattern will be routed to that queue.
- Patterns can contain * (match exactly one word) and # (match zero or more words).

# 4. Headers Exchange

![Headers Exchange](https://www.cloudamqp.com/img/blog/rabbitmq-headers-exchange.svg)

A headers exchange routes messages based on message header values rather than a routing key. The binding can specify headers and values, and an optional match type (all or any).

Use Case: When you want to route messages based on multiple attributes that are best expressed in headers, rather than constructing a complex routing key.

```go
package main

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // 1. Connect to RabbitMQ
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // 2. Create a channel
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // 3. Declare a headers exchange
    exchangeName := "headers_logs"
    err = ch.ExchangeDeclare(
        exchangeName,
        "headers", // type
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare headers exchange")

    // 4. Declare a queue
    queueName := "headers_log_queue"
    q, err := ch.QueueDeclare(
        queueName,
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare queue")

    // 5. Bind the queue to the headers exchange.
    //    We specify the headers to match and the match type via "x-match".
    //    Possible "x-match" values: "all" or "any"
    err = ch.QueueBind(
        q.Name,
        "", // routing key is usually unused in headers exchange
        exchangeName,
        false,
        amqp.Table{
            "x-match": "all",
            "format":  "pdf",
            "type":    "report",
        },
    )
    failOnError(err, "Failed to bind queue to headers exchange")

    // 6. Publish a message with headers
    //    Note: The routing key is ignored by a headers exchange.
    body := "Hello Headers Exchange!"
    err = ch.Publish(
        exchangeName,
        "",
        false,
        false,
        amqp.Publishing{
            Headers: amqp.Table{
                "format": "pdf",
                "type":   "report",
            },
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    failOnError(err, "Failed to publish a message")

    log.Printf("[x] Sent %s with headers format=pdf, type=report", body)

    // 7. Consume messages
    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to register consumer")

    go func() {
        for d := range msgs {
            log.Printf("Received: %s", d.Body)
        }
    }()

    log.Println("[*] Waiting for messages. To exit press CTRL+C")
    select {}
}
```

## How It Works
- Instead of routing keys, headers exchanges look at message headers.
- The queue binding can specify multiple headers.
- "x-match" = "all" means all headers must match. "x-match" = "any" means any of the listed headers can match.

# Summary of Exchange Types

![Some Exchange Types](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*hMLDKDt-SVwAWfOw0sHsqg.gif)

1. Direct Exchange
Routing Mechanism: Exact match on routing key.
Common Use: Selective 1:1 routing.

2. Fanout Exchange
Routing Mechanism: Broadcasts to all bound queues (ignores routing key).
Common Use: One-to-all broadcasting.

3. Topic Exchange
Routing Mechanism: Pattern matching on the routing key (using * and #).
Common Use: Publish/subscribe with wildcard-based routing (logs, events with multiple topics).

4. Headers Exchange
Routing Mechanism: Based on message headers instead of routing keys; supports matching strategies (all, any).

---

Each code sample follows a similar pattern:
1. Connect to RabbitMQ.
2. Create a channel.
3. Declare an exchange of the appropriate type.
4. Declare a queue.
5. Bind the queue to the exchange (with the relevant binding: routing key, pattern, or headers).
6. Publish a message (with either a routing key or headers).
7. Consume messages from the queue.

By changing the exchange type, binding parameters, and publishing parameters, you can adapt the same core code to use any of the four main exchange types in RabbitMQ.

---

There are two more exchange types:

1. Default Exchange     
The default exchange is a pre-declared direct exchange that has no name. It is usually referred by an empty string. If you use default exchange your message is delivered to the queue with a name equal to the routing key of the message. Every queue is automatically bound to the default exchange with a routing key which is the same as the queue name.

2. Dead Letter Exchange     
If there is no matching queue for the message, the message is dropped. RabbitMQ provides an AMQP extension known as the “Dead Letter Exchange”. This exchange which provides the functionality to capture messages that are not deliverable.