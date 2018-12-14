package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type HTTPEchoService struct {
	listen string
}

func (h *HTTPEchoService) Flags(fs *flag.FlagSet) {
	fs.StringVar(&h.listen, "http-addr", "127.0.0.1:8080", "HTTP listen address")
}

func (h *HTTPEchoService) Run() error {
	http.HandleFunc("/", httpHandler)
	log.Println("Listening for HTTP on", h.listen)
	return http.ListenAndServe(h.listen, nil)
}

type HTTPRequestDebug struct {
	Transport struct {
		Protocol string `json:"protocol"`
		Method   string `json:"method"`
		URL      string `json:"url"`
		Host     string `json:"host"`
	} `json:"transport"`
	Remote struct {
		Address string `json:"address"`
	} `json:"remote"`
	Headers http.Header `json:"headers"`
	// TODO: Forms, content body, etc.
}

func newHTTPRequestDebug(r *http.Request) HTTPRequestDebug {
	return HTTPRequestDebug{
		Transport: struct {
			Protocol string `json:"protocol"`
			Method   string `json:"method"`
			URL      string `json:"url"`
			Host     string `json:"host"`
		}{
			Protocol: r.Proto,
			Method:   r.Method,
			URL:      r.URL.String(),
			Host:     r.Host,
		},
		Remote: struct {
			Address string `json:"address"`
		}{
			Address: r.RemoteAddr,
		},
		Headers: r.Header,
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Build up a representation of the HTTP request.
	// 2. Output it as accept requests.
	// TODO: Smarter Accept parsing, it can be much more complicated
	switch r.Header.Get("Accept") {
	case "*/*":
		fallthrough
	case "text/*":
		fallthrough
	case "text/plain":
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(dump)
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		err := enc.Encode(newHTTPRequestDebug(r))
		if err != nil {
			log.Println(err)
		}
	case "text/html":
		// TODO
		http.Error(w, "Not Implemented Yet", http.StatusNotFound)
	default:
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
	}
}
