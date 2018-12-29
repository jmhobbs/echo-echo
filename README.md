# echo-echo

A multi-protocol echo server for debugging.

## Usage

    Usage of echo-echo:
      -config string
          config file (optional)
      -http-addr string
          HTTP listen address (default "127.0.0.1:8080")
      -tcp-addr string
          TCP service listen address (default "127.0.0.1:9000")
      -tcp-delimiter string
          TCP service echo delimiter (one character) (default "\n")

## HTTP Server

    jmhobbs@localhost:~/ $ curl -H 'X-Custom-Header: Value' http://127.0.0.1:8080/
    GET / HTTP/1.1
    Host: 127.0.0.1:8080
    Accept: */*
    User-Agent: curl/7.54.0
    X-Custom-Header: Value

    jmhobbs@localhost:~/ $ curl -H 'X-Custom-Header: Value' -H 'Accept: application/json' http://127.0.0.1:8080/
    {
      "transport": {
        "protocol": "HTTP/1.1",
        "method": "GET",
        "url": "/",
        "host": "127.0.0.1:8080"
      },
      "remote": {
        "address": "127.0.0.1:51325"
      },
      "headers": {
        "Accept": [
          "application/json"
        ],
        "User-Agent": [
          "curl/7.54.0"
        ],
        "X-Custom-Header": [
          "Value"
        ]
      }
    }

## TCP Service

    jmhobbs@localhost:~/ $ telnet 127.0.0.1 9000
    Trying 127.0.0.1...
    Connected to localhost.
    Escape character is '^]'.
    Hello!
    Hello!
    ^]
    telnet> Connection closed.
