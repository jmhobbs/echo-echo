package main

import (
	"flag"
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

	ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	go httpSvc.Run()
	go tcpSvc.Run()

	done := make(chan bool)
	<-done
}

type EchoService interface {
	Flags(*flag.FlagSet)
	Run()
}
