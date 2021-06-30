#go语言学习笔记
##GO的特点
```
go语言不存在隠式转换，因此所有的转换都必须显式说明
gofmt强制统一代码格式
常量使用const定义，用于存储不会改变的数据，声明变量一般使用var关键字
当一个变量被声明时，系统会赋予初始值，int为0，float为0.0，boole为false,string为空字符串，指针为nil
函数不能被重载，Go 语言不支持这项特性的主要原因是函数重载需要进行多余的类型匹配影响性能；没有重载意味着只是一个简单的函数调度。
不要使用 new，永远用 make 来构造 map
像 fmt、os 等这样具有常用功能的内置包在 Go 语言中有 150 个以上，它们被称为标准库，大部分(一些底层的除外)内置于 Go 本身。完整列表可以在 Go Walker 查看(https://gowalker.org/search?q=gorepos)
```
##Go 程序的一般结构
```
先导入包，在完成包的导入时，声明变量、常量和类型的定义
如果存在init函数时，则应对init函数进行定义（这是一个特殊的函数，每个含有该函数的包都应该先执行这个函数）
如果当前的包是main包，则应该先定义main函数
然后定义其他的函数，首先是类型的方法，接着时按照main函数先后的调用顺序来定义相关函数
```