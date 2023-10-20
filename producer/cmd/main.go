package main

import (
	"fmt"
	"time"
	stream "wb/producer/repository"
)

func main() {
	// stream factory?
	conn, err := stream.Connect()
	defer stream.Close(conn)

	if err != nil {
		fmt.Println(err)
	}

	for {
		stream.Publish(conn, []byte("hello from producer"))

		time.Sleep(time.Second * 5)
	}
}
