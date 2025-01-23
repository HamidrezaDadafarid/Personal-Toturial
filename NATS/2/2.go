package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go" // Import the NATS Go client library
)

func main() {
	// Connect to the NATS server (default is nats://localhost:4222)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		// If there's an error during connection, log it and terminate
		log.Fatal(err)
	}
	defer nc.Close() // Ensure the connection is closed when the program exits

	// Define the subject for communication
	subject := "service.request"

	// Start a goroutine to act as the service (listener) that handles requests
	go func() {
		// Subscribe to the subject and define the handler for incoming messages
		nc.Subscribe(subject, func(m *nats.Msg) {
			// Print the received request message
			fmt.Printf("Received request: %s\n", string(m.Data))

			// Prepare a response (this could be dynamic based on the request)
			response := "Here is your response!"

			// Publish the response back to the reply subject included in the message
			err := nc.Publish(m.Reply, []byte(response))
			if err != nil {
				// If there's an error while publishing the reply, log it and terminate
				log.Fatal(err)
			}

			// Log that the response was sent
			fmt.Printf("Sent reply: %s\n", response)
		})

		// Keep the listener running indefinitely by blocking with a `select` statement
		select {}
	}()

	// Give the listener some time to start
	time.Sleep(1 * time.Second)

	// Make a request to the service (subject) with a message and wait for a reply
	msg, err := nc.Request(subject, []byte("Hello, can I get a response?"), 2*time.Second)
	if err != nil {
		// If there's an error (e.g., no reply within timeout), log it and terminate
		log.Fatal(err)
	}

	// Print the received reply message
	fmt.Printf("Received reply: %s\n", string(msg.Data))

	// Flush any buffered messages to the server
	nc.Flush()

	// Check for any errors during the flush
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
}
