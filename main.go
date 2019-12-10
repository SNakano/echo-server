package main

import (
	"flag"
	"net/http"
	"strings"
)

var (
	listenFlag = flag.String("listen", ":8080,:8081", "address and ports to listen")
	textFlag   = flag.String("text", "Hello world!", "text to put on the webpage")
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK!"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(*textFlag))
}

var servers = []*http.ServeMux{}

func main() {
	flag.Parse()
	finish := make(chan bool)

	for _, listen := range strings.Split(*listenFlag, ",") {
		var server = http.NewServeMux()
		server.HandleFunc("/", hello)
		server.HandleFunc("/health", health)
		servers = append(servers, server)
		go http.ListenAndServe(listen, server)
	}
	<-finish
}
