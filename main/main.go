package main

import (
	"../chan"
	"fmt"
)

func main() {
	channel := make(chan int)
	go _chan.Circular1W(channel)
	<-channel
	fmt.Println("main over")
}
