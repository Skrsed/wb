package main

import (
	"fmt"
	"net/http"
	//"wb/producer/repository/stan"
	//stan "github.com/nats-io/stan.go
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from http server changed!")
}

func main() {
	// set up routes
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", homeHandler)

	// TODO: check name conventions
	// clusterID := "test-cluster"
	// clientID := "wb/publisher"

	// connect to stan
	// defer func() {
	// 	sc, _ := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))

	// 	err := sc.Publish("foo", []byte("Hello World"))

	// 	fmt.Errorf("%v", err)
	// 	fmt.Printf("here")
	// }()

	// run listen server
	//log.Fatal(http.ListenAndServe("localhost:80", mux))
}
