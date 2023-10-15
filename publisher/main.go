package main

import (
	"fmt"
	"log"
	"net/http"

	stan "github.com/nats-io/stan.go"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from http server changed!")
}

func main() {
	// set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	// TODO: check name conventions
	clusterID := "cluster"
	clientID := "wb/publisher"

	// connect to stan
	sc, _ := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:1234"))

	sc.Publish("foo", []byte("Hello World"))

	// run listen server
	log.Fatal(http.ListenAndServe("localhost:80", mux))
}
