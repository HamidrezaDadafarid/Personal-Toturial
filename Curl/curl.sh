#!/usr/bin/bash
# Specifies the shell to use for executing the script

# Sets strict error handling:
# -e: Exit immediately if a command exits with a non-zero status
# -u: Treat unset variables as an error and exit immediately
# -o pipefail: Return the exit status of the last command in the pipeline that failed
set -e -u -o pipefail

# Fetches all posts from the placeholder API
curl https://jsonplaceholder.typicode.com/posts

# Fetches the post with ID 3 from the placeholder API
curl https://jsonplaceholder.typicode.com/posts/3

# Fetches the post with ID 21 and includes the HTTP headers in the response
curl -i https://jsonplaceholder.typicode.com/posts/21

# Retrieves only the HTTP headers for post ID 21
curl --head https://jsonplaceholder.typicode.com/posts/21
# Another way to retrieve only HTTP headers for all posts
curl -I https://jsonplaceholder.typicode.com/posts

# Displays the detailed process of the request, including headers and response
curl -v https://jsonplaceholder.typicode.com/posts

# Downloads the response of all posts to a file named output.txt
curl -o output.txt https://jsonplaceholder.typicode.com/posts

# Downloads the response of all posts to a file named outout.txt (identical to -o option)
curl --output output.txt https://jsonplaceholder.typicode.com/posts

# Another way to redirect the output to output.txt
curl https://jsonplaceholder.typicode.com/posts > output.txt

# Runs curl silently (no output to terminal), saves response to output.txt
curl -s -o output.txt https://jsonplaceholder.typicode.com/posts

# Runs curl silently and saves the response to output.json
curl -s https://httpbin.org/get > output.json

# Downloads the file using the same filename as it exists on the server (e.g., 'posts')
curl -O https://jsonplaceholder.typicode.com/posts

# Downloads the file with a limited download speed of 1000 bytes per second
curl -O --limit-rate 1000B https://jsonplaceholder.typicode.com/posts

# Sends a POST request with JSON data to create a new post (the dafault methood if we dont specify it is POST when we are sending data)
curl -d '{"title": "title1", "body": "body1"}' https://jsonplaceholder.typicode.com/posts
# Alternative way to send the same POST request
curl --data '{"title": "title1", "body": "body1"}' -X POST https://jsonplaceholder.typicode.com/posts

# Sends a PUT request to update the title of the post with ID 3
curl -d '{"title": "new-title"}' -X PUT https://jsonplaceholder.typicode.com/posts/3

# Sends a PUT request with JSON data and specifies the content type as application/json
curl -d '{"title": "new-title"}' -X PUT -H 'Content-Type: application/json' https://jsonplaceholder.typicode.com/posts/3

# Sends a DELETE request to delete the post with ID 3
curl -X DELETE https://jsonplaceholder.typicode.com/posts/3

# Authenticates with a username and password (HTTP Basic Auth) to access the resource
curl -u username:password https://example.com

# Sends an Authorization header with a bearer token for authentication
curl -H 'Authorization: Bearer YOUR_TOKEN' https://example.com

# Follows redirects to reach the final destination (useful when the URL redirects to another URL)
curl -L http://google.com

# Follows up to 5 redirects and includes headers in the output for each request/response pair
curl -i -L https://httpbin.org/redirect/5
