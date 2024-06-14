package main

import (
	"github.com/welsonMuniz/fcutils/pkg/rabbitmq"
)

// teste
func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World2!", "amq.direct")
}
