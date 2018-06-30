package main

import (
	"../base"
	"fmt"
	"sync"
)

var once sync.Once
func main() {
	slice := []int{5, 2, 8, 9, 2, 3, 4, 9}
	base.QuickSort(slice, 0, len(slice)-1)
	fmt.Println("quickSort", slice)
}
