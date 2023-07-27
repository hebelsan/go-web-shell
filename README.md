# Go-web-shell

This repository is a simple Golang HTTP server designed to allow for the remote
execution of shell commands.

It was originally designed to automate deployment workflows via webhooks. 

NOTE: This is an insecure approach because the token is not encrypted, and is mainly for demonstration purposes. If you need something like this look into just running your commands with SSH or via Puppet if you need something more robust.

## Security

In order to validate an inbound request, this server checks that a specific header called `token` 
matches what you set in your server configuration. You may also want to firewall the port you
choose only for specfic IP addresses, since this server enables shell access to your
machine. Be careful!

## Installation and Use

Compile the binary with `go build` for your operating system of choice. Then send the binary to your server:

```terminal
$ go build .
```

Run the binary on an open port of your choosing and pass it your API key:

```terminal
$ ./go run main.go --port=5555 --token=20fnoq8yeho8h1y3o
```

You can now pass arbitrary shell commands to the server by POSTing them to the `/cmd` endpoint.

```
$ curl --location --request POST 'http://localhost:5555/cmd' \
--header 'token: 20fnoq8yeho8h1y3o' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "ls",
    "args": [
        "-la"
    ]
}'
```

If the server is already executing a cmd it'll return a 503, if not, you'll
receive a JSON response.