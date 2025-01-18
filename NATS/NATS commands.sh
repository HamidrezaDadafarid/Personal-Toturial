#!/usr/bin/env bash

# Start a NATS server instance on the default address (localhost:4222)
nats-server

# Publishing and subscribing to topics using the NATS CLI
# NATS uses a default address of localhost:4222, unless otherwise specified.
# It is a convention to separate topic tokens using "." (dot). However, other characters like "::", "#", "/" can also be used.

# Publish an event on a topic with two tokens in its name
nats pub msg.test "NATS MESSAGE 1"

# Publish an event on a topic with three tokens in its name
nats pub msg.test1.test2 "NATS MESSAGE 2"

# Subscribe to a topic with exactly two tokens in its name
nats sub "msg.test"

# Subscribe to topics with exactly three tokens in their names
# The asterisk (*) acts as a wildcard for a single token in the topic.
nats sub "msg.*.*"

# Subscribe to topics where the second token can be anything, and the third token must be "test2"
# Example: This matches "msg.xyz.test2" but not "msg.xyz.abc" or "msg.xyz".
nats sub "msg.*.test2"

# Subscribe to a topic ("msg.test") and join the specified queue group ("q1")
# Messages for this subscription will be load-balanced across all subscribers in the same queue group.
nats sub "msg.test" --queue q1

# Subscribe to all topics starting with "msg" (including subtopics at any depth)
# The greater-than symbol (>) is a wildcard for 1 or more tokens in the topic name.
# Example: Matches "msg.test", "msg.test1.test2", "msg.test.abc.xyz", etc.
nats sub "msg.>"

# Subscribe to "msg.test" and save received messages to the specified file
# The "--dump" option specifies the path where received messages are saved.
nats sub "msg.test" --dump "./received-messages"

# Subscribe to a topic with a maximum count of 10 messages and use a different NATS server address
# The "-s" flag specifies a different server address instead of the default localhost:4222.
nats subscribe --count 10 -s another-address:4222 topic

# Subscribe to a topic on a secured server using username and password authentication
# Replace "USER" and "PASSWORD" with the appropriate credentials for the NATS server.
nats sub -s localhost:4222 --user=USER --password=PASSWORD

# Reply to a specific topic with a predefined message
# When a client requests help on "help.please", this subscriber will reply with "OK, I CAN HELP!!!".
nats reply help.please 'OK, I CAN HELP!!!'

# Send a request to the "help.please" topic with a message
# The client sends "I need help!" and expects a reply from a subscriber on the same topic.
nats request help.please 'I need help!'
