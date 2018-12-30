package main

import (
	"flag"
	"log"
	"os"

	"github.com/peterbourgon/ff"
)

func main() {
	fs := flag.NewFlagSet("echo-echo", flag.ExitOnError)
	_ = fs.String("config", "", "config file (optional)")

	httpSvc := &HTTPEchoService{}
	httpSvc.Flags(fs)

	tcpSvc := &TCPEchoService{}
	tcpSvc.Flags(fs)

	dnsSvc := &DNSEchoService{}
	dnsSvc.Flags(fs)

	ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	go httpSvc.Run(log.New(os.Stderr, "http | ", 0))
	go tcpSvc.Run(log.New(os.Stderr, " tcp | ", 0))
	go dnsSvc.Run(log.New(os.Stderr, " dns | ", 0))

	done := make(chan bool)
	<-done
}

type EchoService interface {
	Flags(*flag.FlagSet)
	Run(*log.Logger)
}
