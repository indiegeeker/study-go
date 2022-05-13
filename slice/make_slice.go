package main

import "fmt"

func main() {
	//make 创建qie片
	test := make([]string, 5) //len(test) = 5
	printSlice("test切片", test)

	test1 := make([]string, 0, 5) //len(test1) = 0    cap(test1) = 5
	printSlice("test1切片", test1)

	split_test := test[:2]
	printSlice("切割0,2的数据", split_test)

	split_test1 := test[0:1]
	printSlice("切割0,1的数据", split_test1)
}

func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
