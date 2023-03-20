package main

import "fmt"

func demo(i int) {
	var arr [10]int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 10
}

func main() {
	demo(10)
	fmt.Println("程序继续执行")
}
