package main

import (
	"fmt"

	"golang.org/x/tour/tree"
	//"strconv"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	//fmt.Println(t.Value)
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch0 := make(chan int, 15)
	ch1 := make(chan int, 15)
	Walk(t1, ch0)
	close(ch0)
	Walk(t2, ch1)
	close(ch1)
	l0 := len(ch0)
	l1 := len(ch1)
	if l0 != l1 {
		return false
	}
	for i := 0; i < l0; i++ {
		if <-ch0 != <-ch1 {
			return false
		}
		//fmt.Println("ch0:" + strconv.Itoa(<-ch0))
	}
	//for i := 0; i < l1; i++ {
	//	fmt.Println("ch1:" + strconv.Itoa(<-ch1))
	//}
	return true
}

func main() {
	tree0 := tree.New(2)
	tree1 := tree.New(2)
	ret := Same(tree0, tree1)
	fmt.Println(ret)
}
