//输出前20个斐波那契数列
package main

import fm "fmt"

func main()  {
	result := 0
	for i := 0; i < 20; i++ {
		result = fibonacci(i)
		fm.Printf("第%d个斐波那契数为%d\n",i+1,result)
	}
}

func fibonacci(n int)(res int)  {
	if n<=1 {
		res = 1
	}else{
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return 
}



