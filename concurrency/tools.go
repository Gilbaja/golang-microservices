package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	go helloWorld(c)
	for greeting := range c {
		fmt.Println("Receiving from the channel")
		fmt.Printf("%s World!!!\n", greeting)
		fmt.Println("Received from the channel")
		time.Sleep(1 * time.Second)
	}
	//fmt.Println("Receiving from the channel")
	//fmt.Printf("%s World!!!\n", <-c)
	//fmt.Println("Received from the channel")
}

func helloWorld(c chan string) {
	fmt.Println("Sending 1 to the channel")
	c <- "Hello1"
	fmt.Println("Sending 2 to the channel")
	c <- "Hello2"
	fmt.Println("Sending 3 to the channel")
	c <- "Hello3"
	fmt.Println("Sending 4 to the channel")
	c <- "Hello4"
	close(c)
}
