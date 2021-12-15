package main

import "fmt"

func main()  {
	//打印n*n乘法表
	var n int = 9;
	sub(n);
	fmt.Println();
	gold(n);
	
}

func sub(number int)  {
	for i := 1;i <= number;i++{
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ",j,i,j*i)
		}
		fmt.Println(" ");
	}
}

func gold(number int)  {
	//打印空心金字塔
	/*
		 *   
		* *
	   *****
	   i表示层数
	   j表示打印*
	*/
	for i := 1; i <= number; i++ {
		//在打印*之前打空格
		for k := 1; k <= number-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == number{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}