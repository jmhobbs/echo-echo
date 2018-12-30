package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/miekg/dns"
)

type DNSMap map[string]string

func (m DNSMap) String() string {
	s := []string{}
	for k, v := range m {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, ",")
}

func (m DNSMap) Set(value string) error {
	pair := strings.Split(value, "=")
	m[pair[0]] = pair[1]
	return nil
}

type DNSEchoService struct {
	listen  string
	records DNSMap
	logger  *log.Logger
}

func (d *DNSEchoService) Flags(fs *flag.FlagSet) {
	d.records = make(DNSMap)
	fs.StringVar(&d.listen, "dns-addr", "127.0.0.1:5353", "DNS service listen address")
	fs.Var(&d.records, "dns-map", "DNS domain to IP map in the form of \"domain=ip\". Multiple values accepted.")
}

func (d *DNSEchoService) Run(l *log.Logger) error {
	d.logger = l
	l.Println("Listening on", d.listen)
	dns.HandleFunc(".", d.handle)

	server := &dns.Server{Addr: d.listen, Net: "udp"}
	err := server.ListenAndServe()
	if err != nil {
		l.Println("error:", err)
		return err
	}
	defer server.Shutdown()

	done := make(chan bool)
	<-done

	return nil
}

func (d *DNSEchoService) handle(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range m.Question {
			switch q.Qtype {
			case dns.TypeA:
				ip := d.records[strings.TrimSuffix(q.Name, ".")]
				if ip != "" {
					rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
					if err == nil {
						m.Answer = append(m.Answer, rr)
					}
				}
			default:
				d.logger.Println("Unsupported question:", q.Qtype)
			}
		}
	default:
		d.logger.Println("Unsupported opcode:", r.Opcode)
	}

	w.WriteMsg(m)
}
