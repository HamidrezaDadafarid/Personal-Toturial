#!/usr/bin/bash

# Exit immediately if a command exits with a non-zero status, treat unset variables as an error, and fail if any command in a pipeline fails.
set -e -u -o pipefail

# Make a simple GET request to https://httpbin.org/get and print the response to the console.
curl https://httpbin.org/get

# Make a GET request to https://httpbin.org/get with headers included in the response.
curl -i https://httpbin.org/get

# Make a GET request with verbose output, showing details like request/response headers, connection status, etc.
curl -v https://httpbin.org/get

# Make a silent GET request and save the output to a file named output.json.
curl -s -o output.json https://httpbin.org/get

# Similar to above, silently save the response to output.json using output redirection.
curl -s https://httpbin.org/get > output.json

# Send a POST request with JSON data and display the server's response.
curl -d '{"key1": "value1"}' -X POST https://httpbin.org/post

# Send a POST request with JSON data, specifying the Content-Type as application/json for correct handling on the server.
curl -d '{"key1": "value1"}' -X POST -H 'Content-Type: application/json' https://httpbin.org/post

# Send a GET request with an Authorization header (replace YOUR_TOKEN with a valid token).
curl -H 'Authorization: Bearer YOUR_TOKEN' https://httpbin.org/anything

# Send a GET request that automatically follows up to 5 redirects, expecting an HTML response.
curl -L -X GET https://httpbin.org/redirect/5 -H 'accept: text/html'

# Similar to the above, but with response headers included.
curl -i -L -X GET https://httpbin.org/redirect/5 -H 'accept: text/html'
