package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
)

type TCPEchoService struct {
	listen    string
	delimiter string
}

func (t *TCPEchoService) Flags(fs *flag.FlagSet) {
	fs.StringVar(&t.listen, "tcp-addr", "127.0.0.1:9000", "TCP service listen address")
	fs.StringVar(&t.delimiter, "tcp-delimiter", "\n", "TCP service echo delimiter (one character)")
}

func (t *TCPEchoService) Run(l *log.Logger) error {
	l.Println("Listening on", t.listen)
	s, err := net.Listen("tcp", t.listen)
	if err != nil {
		return err
	}
	for {
		client, err := s.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go t.handleTCP(client)
	}
}

func (t *TCPEchoService) handleTCP(client net.Conn) {
	defer client.Close()

	var delimiter byte = []byte(t.delimiter)[0]

	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes(delimiter)
		if err != nil {
			if err == io.EOF {
				client.Write(line)
			} else {
				log.Println(err)
			}
			return
		}
		client.Write(line)
	}
}
