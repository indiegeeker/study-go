/*
	将函数作为返回值的例子
*/
package main

import "strings"
import fm "fmt"

func main()  {
	//现在，我们可以生成如下函数：
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	//然后调用它们：
	fm.Println(addBmp("file"));
	fm.Println(addJpeg("file"));
}

//一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。下面的函数演示了如何动态返回追加后缀的函数：
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}


