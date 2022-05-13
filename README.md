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



....未完待续