package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	httpAddress := os.Getenv("HTTP_ADDRESS")
	if "" == httpAddress {
		httpAddress = ":8080"
	}

	http.HandleFunc("/", httpHandler)
	log.Println("Listening for HTTP on", httpAddress)
	log.Fatal(http.ListenAndServe(httpAddress, nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(dump)
}
