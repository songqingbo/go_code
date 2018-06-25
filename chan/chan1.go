package _chan

import "fmt"

func Circular1W(channel chan int) {
	for i := 1; i <= 100000; i++ {
		fmt.Println(i)
	}
	channel <- 1
}
