# RabbitMQ Queue Types

RabbitMQ provides multiple queue types, each optimized for different use cases and consistency/trade-off requirements. As of RabbitMQ 3.9+, the main built-in queue types are:
1. Classic queues (the original/“default” queue type)
2. Quorum queues (a modern, replicated queue type based on the Raft consensus protocol)
3. Stream queues (introduced in RabbitMQ 3.9, optimized for high-throughput “streaming” use cases)

Below is an overview of each queue type and how to declare them (in Go) using [github.com/rabbitmq/amqp091-go](github.com/rabbitmq/amqp091-go). We will also touch on additional properties (like lazy queues, exclusive queues, auto-delete, etc.), but those are generally configurations rather than separate “types.”

## 1. Classic Queues

### Overview
- Default type in RabbitMQ (pre-3.8).
- Stores messages in memory (with optional paging to disk).
- Simpler approach to mirroring (HA) via mirroring policies (deprecated in favor of quorum queues, but still supported).
- Suited for lighter or simpler workloads with minimal replication needs, or if you have an older RabbitMQ cluster using mirrored queues.

### Key Features
- Mirroring (HA) can be configured via policies.
- Lazy queues (store messages primarily on disk to reduce memory usage) can also be enabled via policy or an argument.

### Declaration Example
By default, if you do not specify `x-queue-type`, RabbitMQ will create a classic queue. However, to be explicit, you can provide `x-queue-type = "classic"` when declaring:

```go
// Create a classic queue:
qNameClassic := "my_classic_queue"
argsClassic := amqp091.Table{
    "x-queue-type": "classic", 
    // Optional: other arguments can go here
}

_, err = ch.QueueDeclare(
    qNameClassic, // name
    true,         // durable
    false,        // delete when unused
    false,        // exclusive
    false,        // noWait
    argsClassic,  // arguments
)
if err != nil {
    log.Fatalf("Error declaring classic queue: %v", err)
}
```

**Note: If you omit `x-queue-type`, RabbitMQ 3.8+ will default to classic unless you have changed the server default.**

To enable lazy mode for a classic queue, set `"x-queue-mode": "lazy"` in the queue arguments or apply a lazy queue policy.

## 2. Quorum Queues

### Overview
- Introduced in RabbitMQ 3.8 as the recommended replacement for mirrored classic queues.
- Uses the Raft protocol to replicate messages across multiple nodes in a cluster.
- Provides high availability and strong data safety guarantees for consistent message storage.
- Ideal for critical workloads that need to survive node failures without losing messages.

### Key Features
- Leader/Follower replication across nodes.
- Better fault tolerance than mirrored classic queues.
- Consistent replication at the cost of more overhead than a single replica (classic queue without mirroring).

### Declaration Example
To declare a Quorum queue, set x-queue-type = "quorum":

```go
// Create a quorum queue:
qNameQuorum := "my_quorum_queue"
argsQuorum := amqp091.Table{
    "x-queue-type": "quorum",
    // Optional advanced arguments can go here, e.g.:
    // "x-quorum-initial-group-size": int32(5),
}

_, err = ch.QueueDeclare(
    qNameQuorum,
    true,  // durable (required for quorum queues)
    false, // delete when unused
    false, // exclusive
    false, // noWait
    argsQuorum,
)
if err != nil {
    log.Fatalf("Error declaring quorum queue: %v", err)
}
```

Notes:
- Quorum queues must be durable (durable=true).
- If you omit durable=true, the declaration will fail.
- Mirroring and lazy options do not apply to quorum queues (they have their own replication mechanism and store data differently).

## 3. Stream Queues (RabbitMQ 3.9+)

### Overview
- Designed for high-throughput use cases similar to streaming platforms (e.g., Kafka).
- Offers log-based storage, allowing consumers to read from historical positions in the log.
- Scales better for long retention and extremely high message rates.

### Key Features
- Log-based design: messages are stored in an append-only log.
- Consumer offset: a consumer can replay older messages from a given offset/position.
- Retention policies: message retention can be time-based, size-based, or both.
- Does not support all features of classic/quorum queues (e.g., you do not ACK in the same manner; it uses “committed offsets”).

### Declaration Example
To declare a stream queue, set `x-queue-type = "stream"`:

```go
// Create a stream queue:
qNameStream := "my_stream_queue"
argsStream := amqp091.Table{
    "x-queue-type": "stream",
    // Additional stream-specific arguments:
    // - "x-max-length-bytes": int64(1000000000), // 1 GB
    // - "x-stream-max-segment-size-bytes": int64(20000000),
    // - "x-stream-offset": "next", // "first", "last", "next"
}

_, err = ch.QueueDeclare(
    qNameStream,
    true,  // durable recommended
    false, // delete when unused
    false, // exclusive
    false, // noWait
    argsStream,
)
if err != nil {
    log.Fatalf("Error declaring stream queue: %v", err)
}
```

**Important: RabbitMQ must have the stream plugin enabled (since 3.9) to use stream queues. They behave quite differently from classic and quorum queues—especially around how consumption and offsets work.**

## Other Queue Properties
Beyond the core queue types, a few additional properties are worth knowing:

1. Durable vs. Non-Durable
Durable (`durable=true`) means the queue definition (and messages, if persistent) survive RabbitMQ broker restarts.
Non-durable queues will disappear if RabbitMQ restarts.

2. Exclusive Queues
An exclusive queue (`exclusive=true`) is only accessible by the connection that declares it and is automatically deleted when that connection closes. Useful for reply queues in RPC-like patterns or for personal worker queues.

3. Auto-Delete Queues
A queue declared as `autoDelete=true` is automatically removed when the last consumer unsubscribes from it.

4. Arguments
Additional arguments, typically in `amqp091.Table`, control advanced features (e.g., message TTL, dead-letter exchange, etc.).
For example, `x-dead-letter-exchange`, `x-dead-letter-routing-key`, `x-message-ttl`, etc.

## Complete Example Demonstrating Multiple Queue Types
Below is a single Go program that connects to RabbitMQ, declares one queue of each type, and then publishes/consumes a single message to demonstrate:
- A **classic** queue (`my_classic_queue`)
- A **quorum** queue (`my_quorum_queue`)
- A **stream** queue (`my_stream_queue`)

**Note: This example assumes you have RabbitMQ 3.9+ (with the stream plugin enabled if you use the stream queue). If you’re on an older version, you might omit the stream queue part.**

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rabbitmq/amqp091-go"
)

func main() {
    // 1. Connect to RabbitMQ
    conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    // 2. Create a channel
    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    // -------------------------------------------------------------------------
    // Classic Queue Declaration
    // -------------------------------------------------------------------------
    classicQueueName := "my_classic_queue"
    argsClassic := amqp091.Table{
        "x-queue-type": "classic", // optional (classic is the default type)
    }
    _, err = ch.QueueDeclare(
        classicQueueName,
        true,  // durable
        false, // auto-delete
        false, // exclusive
        false, // noWait
        argsClassic,
    )
    if err != nil {
        log.Fatalf("Failed to declare classic queue: %v", err)
    }
    log.Printf("[✔] Declared classic queue %s", classicQueueName)

    // -------------------------------------------------------------------------
    // Quorum Queue Declaration
    // -------------------------------------------------------------------------
    quorumQueueName := "my_quorum_queue"
    argsQuorum := amqp091.Table{
        "x-queue-type": "quorum",
        // Optionally: "x-quorum-initial-group-size": int32(3),
    }
    _, err = ch.QueueDeclare(
        quorumQueueName,
        true,  // durable (required for quorum)
        false, // auto-delete
        false, // exclusive
        false, // noWait
        argsQuorum,
    )
    if err != nil {
        log.Fatalf("Failed to declare quorum queue: %v", err)
    }
    log.Printf("[✔] Declared quorum queue %s", quorumQueueName)

    // -------------------------------------------------------------------------
    // Stream Queue Declaration (RabbitMQ 3.9+ with stream plugin enabled)
    // -------------------------------------------------------------------------
    streamQueueName := "my_stream_queue"
    argsStream := amqp091.Table{
        "x-queue-type": "stream",
        // Example of optional arguments:
        // "x-max-length-bytes": int64(10000000), // 10MB retention
    }
    _, err = ch.QueueDeclare(
        streamQueueName,
        true,  // durable recommended for streams
        false, // auto-delete
        false, // exclusive
        false, // noWait
        argsStream,
    )
    if err != nil {
        log.Fatalf("Failed to declare stream queue (ensure stream plugin is enabled): %v", err)
    }
    log.Printf("[✔] Declared stream queue %s", streamQueueName)

    // -------------------------------------------------------------------------
    // Publish a message to each queue
    // -------------------------------------------------------------------------
    publishBody := "Hello from RabbitMQ!"
    queues := []string{classicQueueName, quorumQueueName, streamQueueName}

    for _, qName := range queues {
        err = ch.PublishWithContext(
            context.Background(), // or context.TODO()
            "",                   // exchange (empty means default direct exchange)
            qName,                // routing key = queue name
            false,                // mandatory
            false,                // immediate
            amqp091.Publishing{
                ContentType: "text/plain",
                Body:        []byte(publishBody),
            },
        )
        if err != nil {
            log.Fatalf("Failed to publish message to %s: %v", qName, err)
        }
        log.Printf("[→] Published message to queue %s", qName)
    }

    // -------------------------------------------------------------------------
    // Consume one message from each queue
    // -------------------------------------------------------------------------
    for _, qName := range queues {
        msgs, err := ch.Consume(
            qName, // queue
            "",    // consumer
            true,  // autoAck
            false, // exclusive
            false, // noLocal
            false, // noWait
            nil,   // args
        )
        if err != nil {
            log.Fatalf("Failed to consume from %s: %v", qName, err)
        }

        // Since we only publish one message, let's just receive one
        msg := <-msgs
        log.Printf("[←] Received from queue %s: %s", qName, msg.Body)
    }

    // Done
    log.Println("Done. Shutting down.")
}
```

### Notes on the Example
- Default Exchange: We used the default (unnamed) direct exchange by specifying an empty string as the exchange name, and we used the queue’s name as the “routing key.” This is a simple way to publish directly into a queue.
- Durability: We set all queues to durable=true. For quorum queues, this is required; for stream queues, it’s highly recommended. For classic queues, you can set them to non-durable if you do not need them to persist after broker restarts.
- Stream Plugin: Declaring a queue with "x-queue-type": "stream" will fail if the stream plugin isn’t enabled or if you’re running an older RabbitMQ version that doesn’t support streams.

## When to Use Each Queue Type

- Classic Queues:
    - Good default if you have simple workloads and do not need multi-node replication (or if you rely on older mirroring features).
    - Familiar and widely supported by existing code.

- Quorum Queues:
    - Recommended over mirrored classic queues for high availability.
    - Provides replication via Raft and stronger data consistency in the event of node failures.

- Stream Queues (3.9+):
    - Ideal for event streaming or high-throughput use cases where you might want to replay messages from a log.
    - Has different semantics than typical AMQP queues (offset-based consumption, custom retention policies).

## Summary
1. Classic: The traditional queue type, simplest to use, optionally mirrored with older RabbitMQ HA features.
2. Quorum: A robust, replicated queue using Raft for higher fault tolerance.
3. Stream: A log-based queue for high-performance streaming with replay.

You declare each queue type by providing the `"x-queue-type"` argument in `amqp091.Table` during queue declaration. The rest of your code (publishing, consuming, etc.) can remain mostly unchanged.

This should cover the main RabbitMQ queue types and illustrate how to declare each using `github.com/rabbitmq/amqp091-go` in Go!