# go指南

## 导出名

- 在go中，如果一个名字以大写字母开头，那么它就是已导出的。在导入一个包时，你只能使用其中已导出的包，未导出的名字不能在包外使用。

​	示例：

```
package main

import (
	"fmt"
	"math"
)

func main() {
	//正确
	fmt.Println(math.Pi)
	//错误
	//fmt.Println(math.pi);
}
```

## 函数

- 函数可以没有参数或者可以接受多个参数

示例：

```
package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

func add(x int, y int) int {
	return x + y
}
```

- 当连续两个或多个函数的已命名形参类型相同时，除最后一个外，其它的都可以省略

示例：

```
package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

func add(x,y int) int {
	return x + y
}
```

- 多值返回

函数可以返回任意数量的返回值

示例：

```
package main

import "fmt"

func main() {
	pre, next := swap("hello", "world")
	fmt.Println(pre, next)
}

func swap(pre, next string) (string, string) {
	return next, pre
}
```

- 命名返回值

go的返回值可被命名，它们会被视作定义在函数顶部的变量

示例：

```
package main

import "fmt"

func main() {
	pre_res, next_res := swap("hello", "world")
	fmt.Println(pre_res, next_res)
}

func swap(pre, next string) (pre_res, next_res string) {
	pre_res = next
	next_res = pre
	return
}
```

##  变量

- var语句用于声明一个变量，跟函数的参数列表一样，类型在最后

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

- 变量的初始化

  变量声明可以包含初始值，每个变量一个

  如果初始值已存在，则可以省略， 变量会从初始值获得类型

  ```
  package main
  
  import "fmt"
  
  var i, j int = 1, 2
  
  func main() {
  	var c, python, java = true, false, "no!"
  	fmt.Println(i, j, c, python, java)
  }
  ```

- 短变量的声明

  在函数中，简洁赋值语句 := 可在类型明确的时候代替var使用

  函数外的每个语句都要以 var func开始，故 := 结构不能在函数外使用

  ```
  package main
  
  import "fmt"
  
  func main() {
  	var i, j int = 1, 2
  	k := 3
  	c, python, java := true, false, "no!"
  
  	fmt.Println(i, j, k, c, python, java)
  }
  ```

## 常量

- 常量的声明与变量类似，只不过使用的是const关键字
- 常量可以是字符串，bool，数字或者字符
- 常量不能使用 := 声明

## 数值常量

- 数值常量是高精度的值

- 一个为制定类型的数值常量由上下文来决定类型

  ```
  package main
  
  import "fmt"
  
  const (
  	// 将 1 左移 100 位来创建一个非常大的数字
  	// 即这个数的二进制是 1 后面跟着 100 个 0
  	Big = 1 << 100
  	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
  	Small = Big >> 99
  )
  
  func needInt(x int) int { return x*10 + 1 }
  func needFloat(x float64) float64 {
  	return x * 0.1
  }
  
  func main() {
  	fmt.Println(needInt(Small))
  	fmt.Println(needFloat(Small))
  	fmt.Println(needFloat(Big))
  }
  ```

## 基本类型

- go的基本类型有

  ```
  string
  boot
  
  int, int8,int 16,int 32,int 64
  uint,uint8,uint16,uint32,uint64
  
  byte  //uint8的别名
  
  rune //int32的别名
  
  float32,float64
  
  complex64 complex128
  ```

- 零值,没有明确初始化值的变量声明会被赋予它们的零值

- 类型转换，在不同类型的变量之间进行转换时需要显式转换

- 类型推导，在声明一个变量而不指定类型时（即使用不带 := 和 var = 表达式语法），变量的类型由右值推导而出

  ```
  package main
  
  import "fmt"
  
  const (
  	// 将 1 左移 100 位来创建一个非常大的数字
  	// 即这个数的二进制是 1 后面跟着 100 个 0
  	Big = 1 << 100
  	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
  	Small = Big >> 99
  )
  
  func needInt(x int) int { return x*10 + 1 }
  func needFloat(x float64) float64 {
  	return x * 0.1
  }
  
  func main() {
  	fmt.Println(needInt(Small))
  	fmt.Println(needFloat(Small))
  	fmt.Println(needFloat(Big))
  }
  ```

# 流程控制

## for

- go只有一种循环：for循环

- 基本的 `for` 循环由三部分组成，它们用分号隔开：

  - 初始化语句：在第一次迭代前执行
  - 条件表达式：在每次迭代前求值
  - 后置语句：在每次迭代的结尾执行

  初始化语句通常为一句短变量声明，该变量声明仅在 `for` 语句的作用域中可见。

  一旦条件表达式的布尔值为 `false`，循环迭代就会终止。

  **注意**：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 `{ }` 则是必须的。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	sum := 0
  	for i := 0; i < 10; i++ {
  		sum += i
  	}
  	fmt.Println(sum)
  }
  ```

- for 是 Go 中的 “while”,你可以去掉分号，因为 C 的 `while` 在 Go 中叫做 `for`。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	sum := 1
  	for sum < 1000 {
  		sum += sum
  	}
  	fmt.Println(sum)
  }
  ```

- 无限循环

  ```
  package main
  
  func main() {
  	for {
  	}
  }
  ```

## if

- Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	count := 10
  	var sum int
  	for i := 0; i < count; i++ {
  		if i%2 == 0 {
  			sum = sum + i
  		}
  	}
  	fmt.Println(sum)
  }
  ```

- if的简短语句，和for一样，if语句可以在条件表达式之前执行一个简单的语句，该语句声明的变量作用域仅在if内

  ```
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  func pow(x, n, lim float64) float64 {
  	if v := math.Pow(x, n); v < lim {
  		return v
  	}
  	return lim
  }
  
  func main() {
  	fmt.Println(
  		pow(3, 2, 10),
  		pow(3, 3, 20),
  	)
  }
  ```

- if和else，在if的简短语句中声明的变量同样可以在任何对应的else快中使用

## switch

- switch是一连串的if-elese的简单写法，它运行第一个值等于条件表达式的case语句

- go只运行选定的case,而非之后所有的case语句

- 除非以fallthrough语句结束，负责分支会自动终止

- go的另一个重要点在于switch的case无需为常量，且取值不必为整数

- switch的case语句从上到下依次执行，直到匹配成功为止

- 没有条件的switch同switch true 一样

 ## defer

- defer语句会将函数推迟到外外层函数返回之后执行

  ```
  package main
  
  import "fmt"
  
  func main() {
  	defer fmt.Println("world")
  
  	fmt.Println("hello")
  }
  ```

- defer的推迟函数会被压入一个栈中，当外层函数返回时，被推迟的函数会按照后进先出的顺序调用

# 指针

- go语言有指针，指针保存了值的内存地址

- 类型 *T指的是指向T类型值的指针，其零值为nil

- &操作符会生成一个指向其操作数的指针

- *操作符表示指针指向的底层值

- go没有指针运算

  ```
  package main
  
  import "fmt"
  
  func main() {
  	i, j := 12, 13
  
  	fmt.Println(&i)
  	fmt.Println(&j)
  
  	p := &i         // 指向 i
  	fmt.Println(*p) // 通过指针读取 i 的值
  	*p = 21
  	fmt.Println(&p) // 通过指针设置 i 的值
  	fmt.Println(i)  // 查看 i 的值
  
  	p = &j         // 指向 j
  	*p = *p / 37   // 通过指针对 j 进行除法运算
  	fmt.Println(j) // 查看 j 的值
  }
  ```

 # 结构体

- 一个结构体（struct）就是一组字段（field）

```
package main

import "fmt"

type test struct {
	x int
	y int
}

func main() {
	fmt.Println(test{11, 22})
}
```

## 结构体字段

- 结构体字段使用点号来访问

```
package main

import "fmt"

type test struct {
	x int
	y int
}

func main() {
	fmt.Println(test{11, 22})

	//点号访问结构体字段
	test := test{23, 25}
	x := test.x
	y := test.y

	fmt.Println(x)
	fmt.Println(y)
}
```

## 结构体指针

- 结构体字段可以通过结构体指针来访问

  如果我们有一个指向结构体的指针 `p`，那么可以通过 `(*p).X` 来访问其字段 `X`。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 `p.X` 就可以。

  ```
  package main
  
  import (
  	"fmt"
  )
  
  type test struct {
  	x int
  	y int
  }
  
  func main() {
  	fmt.Println(test{11, 22})
  
  	//点号访问结构体字段
  	test := test{23, 25}
  	x := test.x
  	y := test.y
  
  	fmt.Println(x)
  	fmt.Println(y)
  
  	//结构体指针访问结构体字段
  	p_test := &test
  	p_test.x = 1e9
  	fmt.Println(p_test)
  }
  ```

## 结构体文法

- 结构体文法通过直接列出字段的值来新分配一个结构体。

- 使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）

- 特殊的前缀 `&` 返回一个指向结构体的指针。

  ```
  package main
  
  import (
  	"fmt"
  )
  
  type test struct {
  	x int
  	y int
  }
  
  var (
  	//结构体文法通过直接列出字段的值来新分配一个结构体。
  	test1      = test{21, 22}
  	test_empty = test{}
  	//使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）
  	test2 = test{x: 2}
  	//特殊的前缀 `&` 返回一个指向结构体的指针。
  	p_test = &test{25, 26}
  )
  
  func main() {
  	fmt.Println(test1, test_empty, test2, p_test)
  }
  ```

# 数组

- 类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

  表达式

  ```
  var a [10]int
  ```

  会将变量 `a` 声明为拥有 10 个整数的数组。

  数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组

  ```
  package main
  
  import "fmt"
  
  func main() {
  	var num [10]string
  
  	num[0] = "hello"
  	num[1] = "world"
  
  	fmt.Println(num)
  
  	primes := [6]int{2, 3, 5, 7, 11, 13}
  	fmt.Println(primes)
  }
  ```

# 切片

## 每个切片数组的大小都是固定的。

而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

```
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```
a[low : high]
```

它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```
a[1:4]
```

## 切片就像数组的引用

切片并不存储任何数据，它只是描述底层数组中的一段

更改切片的元素会修改底层数组中对应的元素

与它共享底层数组的切片都会观测到这些修改

```
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

## 切片文法

```
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

切片文法类似于没有长度的数组文法。

这是一个数组文法：

```
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```
[]bool{true, true, false}
```

## 切片的默认行为	

- 在进行切片时，你可以利用它的默认行为来忽略上下界。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	s := []int{2, 3, 5, 7, 11, 13}
  
  	s = s[1:4]
  	fmt.Println(s)
  
  	s = s[:2]
  	fmt.Println(s)
  
  	s = s[1:]
  	fmt.Println(s)
  }
  ```

  切片下界的默认值为 0，上界则是该切片的长度。

  对于数组

  ```
  var a [10]int
  ```

  来说，以下切片是等价的：

  ```
  a[0:10]
  a[:10]
  a[0:]
  a[:]
  ```

## 切片的长度与容量

- 切片拥有 **长度** 和 **容量**。

- 切片的长度就是它所包含的元素个数。

- 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

- 切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

- 你可以通过重新切片来扩展一个切片，给它提供足够的容量。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	s := []int{2, 3, 5, 7, 11, 13}
  	printSlice(s)
  
  	// 截取切片使其长度为 0
  	s = s[:0]
  	printSlice(s)
  
  	// 拓展其长度
  	s = s[:4]
  	printSlice(s)
  
  	// 舍弃前两个值
  	s = s[2:]
  	printSlice(s)
  }
  
  func printSlice(s []int) {
  	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
  }
  ```

## 切片的零值是nil

- nil 切片的长度和容量都为0且没有底层数组

  ```
  //零值切片  nil
  var test []int
  printSlice(test)
  if test == nil {
  	fmt.Println("nil!")
  }
  ```

## 用make创建切片

- 切片可以使用make的方式创建，这也是创建动态数组的方式

- `make` 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

  ```
  a := make([]int, 5)  // len(a)=5
  ```

  要指定它的容量，需向 `make` 传入第三个参数：

  ```
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
  
  b = b[:cap(b)] // len(b)=5, cap(b)=5
  b = b[1:]      // len(b)=4, cap(b)=4
  ```

- 示例：

  ```
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
  ```

## 切片的切片

- 切片可包含任何类型，甚至包括其他的切片

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

## 向切片追加新元素

- go提供了内建的append函数为切片追加新的元素

  ```
  func append(s []T, vs ...T) []T
  ```

  `append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片，其余类型为 `T` 的值将会追加到该切片的末尾。

  `append` 的结果是一个包含原切片所有元素加上新添加元素的切片。

  当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	//为切片追加新的元素
  	// var test []string
  	var test = make([]string, 5)
  	//使用append追加新元素
  	printSlice(test)
  
  	test = append(test, "zhangan", "lisi") //append之后，切片会按需增长
  	printSlice(test)
  }
  
  func printSlice(s []string) {
  	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
  }
  ```

## Range

- for循环的range形式可遍历切片或映射

- 当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

- 可以将下标或值赋予_来忽略它，若你只需要索引忽略第二个即可

  ```
  package main
  
  import "fmt"
  
  func main() {
  	//for循环的range形式可遍历切片和映射
  	var test []int
  	test = append(test, 1, 2, 3, 4, 5, 6, 6, 35, 43, 23, 1)
  	fmt.Println(test)
  
  	for _, v := range test {
  		fmt.Printf("2*%d = %d\n", v, 2*v)
  	}
  }
  ```

# 映射

- 将键映射到值

- 映射的零值为nil,nil映射既没有键，也不能添加键

- make函数会返回给定类型的映射，并将其初始化备用

  ```
  package main
  
  import "fmt"
  
  type test_map struct {
  	Lat, Long float64
  }
  
  var t_map map[string]test_map
  
  func main() {
  	t_map = make(map[string]test_map)
  	t_map["Bell Labs"] = test_map{
  		40.68433, -74.39967,
  	}
  	fmt.Println(t_map["Bell Labs"])
  }
  ```

## 映射的文法

- 映射的文法与结构体相似，不过必须有键名

  ```
  package main
  
  import "fmt"
  
  type Vertex struct {
  	Lat, Long float64
  }
  
  var m = map[string]Vertex{
  	"Bell Labs": Vertex{
  		40.68433, -74.39967,
  	},
  	"Google": Vertex{
  		37.42202, -122.08408,
  	},
  }
  
  func main() {
  	fmt.Println(m)
  }
  ```

- 若顶级类型只是一个类型名，你可以在文法的元素中省略它。

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

## 修改映射

- 在映射m中插入或者修改元素

  ```
  m[key] = elem
  ```

- 获取元素

  ```
  elem = m[key]
  ```

- 删除元素

  ```
  delete(m,key)
  ```

- 通过双赋值检测某个键是否存在：

  ```
  elem, ok = m[key]
  ```

  若 `key` 在 `m` 中，`ok` 为 `true` ；否则，`ok` 为 `false`。

  若 `key` 不在映射中，那么 `elem` 是该映射元素类型的零值。

  同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

  **注** ：若 `elem` 或 `ok` 还未声明，你可以使用短变量声明：

- 示例

  ```
  package main
  
  import "fmt"
  
  func main() {
  	m := make(map[string]int)
  
  	m["Answer"] = 42
  	fmt.Println("The value:", m["Answer"])
  
  	m["Answer"] = 48
  	fmt.Println("The value:", m["Answer"])
  
  	delete(m, "Answer")
  	fmt.Println("The value:", m["Answer"])
  
  	v, ok := m["Answer"]
  	fmt.Println("The value:", v, "Present?", ok)
  }
  ```

## 映射练习

```
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words_count := make(map[string]int)
	string_array := strings.Fields(s)
	for i := 0; i < len(string_array); i++ {
		words_count[string_array[i]] += 1
	}
	return words_count
}

func main() {
	wc.Test(WordCount)
}
```

# 函数

## 函数也是值，他们可以像其它值一样传递

## 函数值可以用作函数的参数或返回值

```
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

## 函数的闭包

- Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

  例如，函数 `adder` 返回一个闭包。每个闭包都被绑定在其各自的 `sum` 变量上。

  ```
  package main
  
  import "fmt"
  
  func adder() func(int) int {
  	sum := 0
  	return func(x int) int {
  		sum += x
  		return sum
  	}
  }
  
  func main() {
  	pos, neg := adder(), adder()
  	for i := 0; i < 10; i++ {
  		fmt.Println(
  			pos(i),
  			neg(-2*i),
  		)
  	}
  }
  ```

- 练习：斐波那契数列使用闭包的形式实现

  ```
  package main
   
  import "fmt"
   
  // 返回一个“返回int的函数”
  func fibonacci() func() int {
  	a1, a2 := 0, 1
  	return func() int {
  		res := a1
  		a1, a2 = a2, a1+a2
  		return res
  	}
  }
   
  func main() {
  	f := fibonacci()
  	for i := 0; i < 10; i++ {
  		fmt.Println(f())
  	}
  }
  ```

# 方法和接口

## 方法

- go没有类，不过你可以为结构体类型定义方法

- 方法就是一类带特殊的接受者的参数的函数

- 方法接收者在它自己的参数列表内，位于func关键字和方法名之间

  ```
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Vertex struct {
  	X, Y float64
  }
  
  func (v Vertex) Abs() float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
  	v := Vertex{3, 4}
  	fmt.Println(v.Abs())
  }
  
  ```

### 方法即函数

- 方法只是个带接收者参数的函数

### 指针接收者

- 你可以为指针接收者声明方法。

  这意味着对于某类型 `T`，接收者的类型可以用 `*T` 的文法。（此外，`T` 不能是像 `*int` 这样的指针。

  ```
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Vertex struct {
  	X, Y float64
  }
  
  func (v Vertex) Abs() float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func (v *Vertex) Scale(f float64) {
  	v.X = v.X * f
  	v.Y = v.Y * f
  }
  
  func main() {
  	v := Vertex{3, 4}
  	v.Scale(10)
  	fmt.Println(v.Abs())
  }
  ```

### 指针与函数

- 现在我们要把 `Abs` 和 `Scale` 方法重写为函数。

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

### 方法与指针重定向

- 比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

  ```
  var v Vertex
  ScaleFunc(v, 5)  // 编译错误！
  ScaleFunc(&v, 5) // OK
  ```

  而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

  ```
  var v Vertex
  v.Scale(5)  // OK
  p := &v
  p.Scale(10) // OK
  ```

  对于语句 `v.Scale(5)`，即便 `v` 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 `Scale` 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)`。

### 选择值或指针作为接收者

- 使用指针接收者的原因有二：

  首先，方法能够修改其接收者指向的值。

  其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

  在本例中，`Scale` 和 `Abs` 接收者的类型为 `*Vertex`，即便 `Abs` 并不需要修改其接收者。

  通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。

  ```
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Vertex struct {
  	X, Y float64
  }
  
  func (v *Vertex) Scale(f float64) {
  	v.X = v.X * f
  	v.Y = v.Y * f
  }
  
  func (v *Vertex) Abs() float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
  	v := &Vertex{3, 4}
  	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
  	v.Scale(5)
  	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
  }
  ```

## 接口

### 接口类型

接口类型是由一组方法签名定义的集合

```
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

### 接口与隐式实现

- 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。

  隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

  因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

  ```
  package main
  
  import "fmt"
  
  type I interface {
  	M()
  }
  
  type T struct {
  	S string
  }
  
  // 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
  func (t T) M() {
  	fmt.Println(t.S)
  }
  
  func main() {
  	var i I = T{"hello"}
  	i.M()
  }
  ```

### 接口值

- 接口也是值，可以像其他值一样传递

- ‘接口值可以用作函数的参数和返回值

- 在内部，接口值可以看做包含值和具体类型的元组：

  ```
  (value, type)
  ```

  接口值保存了一个具体底层类型的具体值。

  接口值调用方法时会执行其底层类型的同名方法。

  ```
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type I interface {
  	M()
  }
  
  type T struct {
  	S string
  }
  
  func (t *T) M() {
  	fmt.Println(t.S)
  }
  
  type F float64
  
  func (f F) M() {
  	fmt.Println(f)
  }
  
  func main() {
  	var i I
  
  	i = &T{"Hello"}
  	describe(i)
  	i.M()
  
  	i = F(math.Pi)
  	describe(i)
  	i.M()
  }
  
  func describe(i I) {
  	fmt.Printf("(%v, %T)\n", i, i)
  }
  ```

### 底层值为 nil 的接口值

- 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

  在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 `M` 方法）。

  **注意:** 保存了 nil 具体值的接口其自身并不为 nil。

```
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### nil 接口值

- nil接口值即不保存值也不保存具体类型
- 为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 **具体** 方法的类型

```
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### 空接口

- 指定了零个方法的接口值被称为空接口
- 空接口可以保存任何类型的值 （因为每个类型都实现了至少零个方法）
- 空接口被用来处理未知类型的值。例如，`fmt.Print` 可接受类型为 `interface{}` 的任意数量的参数。

### 类型断言

- **类型断言** 提供了访问接口值底层具体值的方式。

  ```
  t := i.(T)
  ```

  该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 `T` 的值赋予变量 `t`。

  若 `i` 并未保存 `T` 类型的值，该语句就会触发一个恐慌。

  为了 **判断** 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

  ```
  t, ok := i.(T)
  ```

  若 `i` 保存了一个 `T`，那么 `t` 将会是其底层值，而 `ok` 为 `true`。

  否则，`ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生恐慌。

  请注意这种语法和读取一个映射时的相同之处。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	var i interface{} = "hello"
  
  	s := i.(string)
  	fmt.Println(s)
  
  	s, ok := i.(string)
  	fmt.Println(s, ok)
  
  	f, ok := i.(float64)
  	fmt.Println(f, ok)
  
  	f = i.(float64) // 报错(panic)
  	fmt.Println(f)
  }
  ```

### 类型选择

- **类型选择** 是一种按顺序从几个类型断言中选择分支的结构。

  类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

  ```
  switch v := i.(type) {
  case T:
      // v 的类型为 T
  case S:
      // v 的类型为 S
  default:
      // 没有匹配，v 与 i 的类型相同
  }
  ```

  类型选择中的声明与类型断言 `i.(T)` 的语法相同，只是具体类型 `T` 被替换成了关键字 `type`。

  此选择语句判断接口值 `i` 保存的值类型是 `T` 还是 `S`。在 `T` 或 `S` 的情况下，变量 `v` 会分别按 `T` 或 `S` 类型保存 `i` 拥有的值。在默认（即没有匹配）的情况下，变量 `v` 与 `i` 的接口类型和值相同。

  ```
  package main
  
  import "fmt"
  
  func do(i interface{}) {
  	switch v := i.(type) {
  	case int:
  		fmt.Printf("Twice %v is %v\n", v, v*2)
  	case string:
  		fmt.Printf("%q is %v bytes long\n", v, len(v))
  	default:
  		fmt.Printf("I don't know about type %T!\n", v)
  	}
  }
  
  func main() {
  	do(21)
  	do("hello")
  	do(true)
  }
  ```

### Stringer

- [`fmt`](https://go-zh.org/pkg/fmt/) 包中定义的 [`Stringer`](https://go-zh.org/pkg/fmt/#Stringer) 是最普遍的接口之一。

  ```
  type Stringer interface {
      String() string
  }
  ```

  `Stringer` 是一个可以用字符串描述自己的类型。`fmt` 包（还有很多包）都通过此接口来打印值。

  ```
  package main
  
  import "fmt"
  
  type Person struct {
  	Name string
  	Age  int
  }
  
  func (p Person) String() string {
  	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
  }
  
  func main() {
  	a := Person{"Arthur Dent", 42}
  	z := Person{"Zaphod Beeblebrox", 9001}
  	fmt.Println(a, z)
  }
  ```

## 错误

- Go 程序使用 `error` 值来表示错误状态。

  与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

  ```
  type error interface {
      Error() string
  }
  ```

  （与 `fmt.Stringer` 类似，`fmt` 包在打印值时也会满足 `error`。）

  通常函数会返回一个 `error` 值，调用的它的代码应当判断这个错误是否等于 `nil` 来进行错误处理。

  ```
  i, err := strconv.Atoi("42")
  if err != nil {
      fmt.Printf("couldn't convert number: %v\n", err)
      return
  }
  fmt.Println("Converted integer:", i)
  ```

  `error` 为 nil 时表示成功；非 nil 的 `error` 表示失败。

  ```
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  type MyError struct {
  	When time.Time
  	What string
  }
  
  func (e *MyError) Error() string {
  	return fmt.Sprintf("at %v, %s",
  		e.When, e.What)
  }
  
  func run() error {
  	return &MyError{
  		time.Now(),
  		"it didn't work",
  	}
  }
  
  func main() {
  	if err := run(); err != nil {
  		fmt.Println(err)
  	}
  }
  ```

## Reader

- `io` 包指定了 `io.Reader` 接口，它表示从数据流的末尾进行读取。

  Go 标准库包含了该接口的[许多实现](https://go-zh.org/search?q=Read#Global)，包括文件、网络连接、压缩和加密等等。

  `io.Reader` 接口有一个 `Read` 方法：

  ```
  func (T) Read(b []byte) (n int, err error)
  ```

  `Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

  示例代码创建了一个 [`strings.Reader`](https://go-zh.org/pkg/strings/#Reader) 并以每次 8 字节的速度读取它的输出。

  ```
  package main
  
  import (
  	"fmt"
  	"io"
  	"strings"
  )
  
  func main() {
  	r := strings.NewReader("Hello, Reader!")
  
  	b := make([]byte, 8)
  	for {
  		n, err := r.Read(b)
  		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
  		fmt.Printf("b[:n] = %q\n", b[:n])
  		if err == io.EOF {
  			break
  		}
  	}
  }
  ```

### rot13Reader

- 有种常见的模式是一个 [`io.Reader`](https://go-zh.org/pkg/io/#Reader) 包装另一个 `io.Reader`，然后通过某种方式修改其数据流。

  例如，[`gzip.NewReader`](https://go-zh.org/pkg/compress/gzip/#NewReader) 函数接受一个 `io.Reader`（已压缩的数据流）并返回一个同样实现了 `io.Reader` 的 `*gzip.Reader`（解压后的数据流）。

  编写一个实现了 `io.Reader` 并从另一个 `io.Reader` 中读取数据的 `rot13Reader`，通过应用 [rot13](http://en.wikipedia.org/wiki/ROT13) 代换密码对数据流进行修改。

  `rot13Reader` 类型已经提供。实现 `Read` 方法以满足 `io.Reader`。

  ```
  package main
  
  import (
  	"io"
  	"os"
  	"strings"
  )
  
  type rot13Reader struct {
  	r io.Reader
  }
  // 转换byte  前进13位/后退13位
  func rot13(b byte) byte {
  	switch {
  	case 'A' <= b && b <= 'M':
  		b = b + 13
  	case 'M' < b && b <= 'Z':
  		b = b - 13
  	case 'a' <= b && b <= 'm':
  		b = b + 13
  	case 'm' < b && b <= 'z':
  		b = b - 13
  	}
  	return b
  }
  // 重写Read方法
  func (mr rot13Reader) Read(b []byte) (int, error) {
  	n, e := mr.r.Read(b)
  	for i := 0; i < n; i++ {
  		b[i] = rot13(b[i])
  	}
  	return n, e
  }
  func main() {
  	s := strings.NewReader("Lbh penpxrq gur pbqr!")
  	r := rot13Reader{s}
  	io.Copy(os.Stdout, &r)
  }
  ```

## 图像

- [`image`](https://go-zh.org/pkg/image/#Image) 包定义了 `Image` 接口：

  ```
  package image
  
  type Image interface {
      ColorModel() color.Model
      Bounds() Rectangle
      At(x, y int) color.Color
  }
  ```

  **注意:** `Bounds` 方法的返回值 `Rectangle` 实际上是一个 [`image.Rectangle`](https://go-zh.org/pkg/image/#Rectangle)，它在 `image` 包中声明。

  （请参阅[文档](https://go-zh.org/pkg/image/#Image)了解全部信息。）

  `color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由 [`image/color`](https://go-zh.org/pkg/image/color/) 包定义。

  ```
  package main
  
  import (
  	"fmt"
  	"image"
  )
  
  func main() {
  	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
  	fmt.Println(m.Bounds())
  	fmt.Println(m.At(0, 0).RGBA())
  }
  ```

### 图像练习

- 还记得之前编写的[图片生成器](https://tour.go-zh.org/moretypes/18) 吗？我们再来编写另外一个，不过这次它将会返回一个 `image.Image` 的实现而非一个数据切片。

  定义你自己的 `Image` 类型，实现[必要的方法](https://go-zh.org/pkg/image/#Image)并调用 `pic.ShowImage`。

  `Bounds` 应当返回一个 `image.Rectangle` ，例如 `image.Rect(0, 0, w, h)`。

  `ColorModel` 应当返回 `color.RGBAModel`。

  At` 应当返回一个颜色。上一个图片生成器的值 `v` 对应于此次的 `color.RGBA{v, v, 255, 255}

```
package main

import (
    "golang.org/x/tour/pic"
    "image/color"
    "image"
)

type Image struct{}  //新建一个Image结构体

func (i Image) ColorModel() color.Model{  //实现Image包中颜色模式的方法
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle{  //实现Image包中生成图片边界的方法
    return image.Rect(0,0,200,200)
}

func (i Image) At(x,y int) color.Color{  //实现Image包中生成图像某个点的方法
    return color.RGBA{uint8(x),uint8(y),uint8(255),uint8(255)}
}

func main() {
    m := Image{}
    pic.ShowImage(m)  //调用
}
```

## 并发

### go程

- Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

  ```
  go f(x, y, z)
  ```

  会启动一个新的 Go 程并执行

  ```
  f(x, y, z)
  ```

  `f`, `x`, `y` 和 `z` 的求值发生在当前的 Go 程中，而 `f` 的执行发生在新的 Go 程中。

  Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。[`sync`](https://go-zh.org/pkg/sync/) 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法。

  ```
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  func say(s string) {
  	for i := 0; i < 5; i++ {
  		time.Sleep(100 * time.Millisecond)
  		fmt.Println(s)
  	}
  }
  
  func main() {
  	go say("world")
  	say("hello")
  }
  ```

### 信道

- 信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值。

  ```
  ch <- v    // 将 v 发送至信道 ch。
  v := <-ch  // 从 ch 接收值并赋予 v。
  ```

  （“箭头”就是数据流的方向。）

  和映射与切片一样，信道在使用前必须创建：

  ```
  ch := make(chan int)
  ```

  默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

  以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

  ```
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
  ```

### 带缓冲的信道

- 信道可以是 *带缓冲的*。将缓冲长度作为第二个参数提供给 `make` 来初始化一个带缓冲的信道：

  ```
  ch := make(chan int, 100)
  ```

  仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

  修改示例填满缓冲区，然后看看会发生什么。

  ```
  package main
  
  import "fmt"
  
  func main() {
  	ch := make(chan int, 3)
  	ch <- 1
  	ch <- 2
  	ch <- 3
  	fmt.Println(<-ch)
  	fmt.Println(<-ch)
  	fmt.Println(<-ch)
  }
  ```

### range和close

- 发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

  ```
  v, ok := <-ch
  ```

  之后 `ok` 会被设置为 `false`。

  循环 `for i := range c` 会不断从信道接收值，直到它被关闭。

  *注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

  *还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 `range` 循环。

  ```
  package main
  
  import (
  	"fmt"
  )
  
  func fibonacci(n int, c chan int) {
  	x, y := 0, 1
  	for i := 0; i < n; i++ {
  		c <- x
  		x, y = y, x+y
  	}
  	close(c)
  }
  
  func main() {
  	c := make(chan int, 10)
  	go fibonacci(cap(c), c)
  	for i := range c {
  		fmt.Println(i)
  	}
  }
  ```

### select语句

- `select` 语句使一个 Go 程可以等待多个通信操作。

  `select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

- 默认选择

  当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

  为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

  ```
  select {
  case i := <-c:
      // 使用 i
  default:
      // 从 c 中接收会阻塞时执行
  }
  ```
