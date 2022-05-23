package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	//go程前半部分
	go sum(s[:len(s)/2], c)
	//go程后半部分
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x+y)
}
