package main

// import "fmt"
import fm "fmt" //import alias

func main() {
   /*
   go 使用var关键字声明变量
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)


   // 声明一个变量并初始化
   var a = "RUNOOB"
   fmt.Println(a)

   // 没有初始化就为零值
   var b int
   fmt.Println(b)

   // bool 零值为 false
   var c bool
   fmt.Println(c)

   var i int
   var f float64
   var b bool
   var s string
   fmt.Printf("%v %v %v %q\n", i, f, b, s)
   */
   
   // := 省略var声明一个变量,格式为 defined_name := defined_value
   // name := "golang"
   // fmt.Printf(name)

   //多变量声明
   //g, h := 123, "hello"这种不带声明格式的只能在函数体中出现
   g, h := 123, "hello"
   fm.Println(g,h) //使用别名fm导入fmt包

}