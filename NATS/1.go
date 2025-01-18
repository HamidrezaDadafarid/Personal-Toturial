package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Importing the NATS Go client package

func main() {
	// Connect to the NATS server (default is nats://localhost:4222)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		// If there is an error while connecting, log it and exit
		log.Fatal(err)
	}
	defer nc.Close() // Ensure that the connection is closed when the function exits

	// Define the subject and message
	subject := "updates"      // The subject to which messages will be published and subscribed
	message := "Hello, NATS!" // The message to be sent to subscribers

	// Subscribe to the 'updates' subject
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		// This callback function is called whenever a message is received on the 'updates' subject
		// 'm' is the message object which contains information like the subject and data
		fmt.Printf("Message Subject: %s\n", m.Subject)   // Print the subject of the message
		fmt.Printf("Message Data: %s\n", string(m.Data)) // Print the message content (converted from byte slice to string)
	})
	if err != nil {
		// If there is an error while subscribing, log it and exit
		log.Fatal(err)
	}

	// Publish a message to the 'updates' subject
	err = nc.Publish(subject, []byte(message)) // Convert the message to a byte slice and publish
	if err != nil {
		// If there is an error while publishing, log it and exit
		log.Fatal(err)
	}
	fmt.Printf("Published message: %s\n", message) // Print the message that was published

	// Flush the connection to ensure that the message is sent
	nc.Flush()

	// Check if there were any errors after flushing the connection
	if err := nc.LastError(); err != nil {
		// If there is an error, log it and exit
		log.Fatal(err)
	}

	// Sleep for 1 second to allow time for the message to be received and processed
	time.Sleep(1 * time.Second)
}

/*
Subscription:
nc.Subscribe(subject, func(m *nats.Msg) {...}): This line sets up the subscription to the updates subject before the message is published.
This means that as soon as the program starts and the NATS connection is established, it begins listening for messages on the updates subject.
The callback function passed to Subscribe will be triggered whenever a message is received on that subject.
*/

/*
Publication:
nc.Publish(subject, []byte(message)): This sends the message to the updates subject after the subscription is set up.
Since the subscription was already established, the message sent by Publish will trigger the callback function, and the subscriber will print out the message details.
*/

/*
Why Subscribe Before Publish?
Avoid Missing Messages: If you publish a message before setting up the subscription, there's a risk that the message could be sent before the subscriber has had a chance to start listening. This would cause the subscriber to miss that message.
Order of Operations: By subscribing first, you ensure that as soon as a message is published, the subscriber is ready to process it.
*/

/*
What's Actually Happening Internally:
NATS has a non-blocking behavior, meaning that it doesn't wait for a subscription to be acknowledged or "ready" before the message is published.
The subscription registers a handler (callback) to be executed when a message arrives. So, when Publish is called, even though the subscription hasn't yet received the message, the message will still be correctly routed to the subscriber once it is published.
*/

/*
Flush Method

Purpose:
Ensures that all buffered data (like published messages) is sent to the NATS server immediately.

How it Works:
When you publish a message (nc.Publish(...)), the NATS client library may buffer the message before sending it to the server. This buffering optimizes performance but introduces a slight delay.
Calling Flush() ensures that the client sends all buffered messages and waits for the server to acknowledge that it received them.

Why It’s Important:
Without calling Flush(), there might be situations where your program exits before the message is sent to the server.
Example: If your program publishes a message and terminates immediately, the message might not be delivered because the client didn't have time to send it.

Key Points:
It blocks until the server acknowledges all sent messages.
It's a safeguard to make sure your messages actually leave the client.
*/

/*
LastError Method

Purpose:
Retrieves the last error encountered during communication with the NATS server.

How it Works:
During operations like Connect, Publish, or Flush, if something goes wrong, the NATS client library logs the error internally.
LastError() provides a way to check if any errors occurred since the connection was established or since the last operation.

Why It’s Useful:
It's a way to ensure that the Flush() operation was successful. For example:

nc.Flush()
if err := nc.LastError(); err != nil {
    log.Fatal(err)
}

If Flush() fails (e.g., due to a network issue or server unavailability), LastError() captures that failure so you can handle it appropriately.

Key Points:
If there's no error, LastError() returns nil.
It's especially useful for debugging and ensuring your messages are properly sent and acknowledged.
*/
