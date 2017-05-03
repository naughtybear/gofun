package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	siginals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-siginals:
		fmt.Println("received siginal", sig)
	default:
		fmt.Println("no activity")
	}
}
