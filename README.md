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

## DNS Service

The DNS service isn't exactly an echo (since that's not really how it works) but instead you specify domain to IP mappings
and it will respond to A record queries with that information.

Note that on a Mac, you will need to either kill mDNSResponder (who uses Bonjour anyway ;) or use a different port `-dns-addr`.

    --[ Terminal A ]---------------------------------------------------------------------------------------
    jmhobbs@localhost:~/ $ echo-echo -dns-map example.com=127.0.0.2 -dns-map google.com=10.0.0.1
    2018/12/29 21:33:13 Listening for TCP on 127.0.0.1:9000
    2018/12/29 21:33:13 Listening for DNS on 127.0.0.1:5353
    2018/12/29 21:33:13 Listening for HTTP on 127.0.0.1:8080

    --[ Terminal B ]---------------------------------------------------------------------------------------
    jmhobbs@localhost:~/ $ dig -p 5353 @localhost example.com

    ; <<>> DiG 9.10.6 <<>> -p 5353 @localhost example.com
    ; (2 servers found)
    ;; global options: +cmd
    ;; Got answer:
    ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 57302
    ;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
    ;; WARNING: recursion requested but not available

    ;; QUESTION SECTION:
    ;example.com.			IN	A

    ;; ANSWER SECTION:
    example.com.		3600	IN	A	127.0.0.2

    ;; Query time: 0 msec
    ;; SERVER: 127.0.0.1#5353(127.0.0.1)
    ;; WHEN: Sat Dec 29 21:33:20 CST 2018
    ;; MSG SIZE  rcvd: 56
