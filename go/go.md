# Module

## 模块管理

```bash
# 初始化模块
go mod init example.com/myproject

# 整理依赖：添加缺失的依赖/移除未使用的依赖
go mod tidy

# 下载依赖
go mod download

# 验证依赖
go mod verify

# 查看依赖图
go mod graph

# 设置 GOPROXY
export GOPROXY=https://goproxy.cn,direct

# 获取最新版本
go get github.com/gin-gonic/gin

# 获取指定版本
go get github.com/gin-gonic/gin@v1.9.0

# 获取特定 commit
go get github.com/gin-gonic/gin@abc1234

# 获取特定分支
go get github.com/gin-gonic/gin@master

# 更新到最新次要版本
go get -u github.com/gin-gonic/gin

# 更新到最新补丁版本
go get -u=patch github.com/gin-gonic/gin

# 更新所有依赖
go get -u ./...

# 移除依赖/运行 `go mod tidy` 清理
go get github.com/gin-gonic/gin@none
```

## 编译/运行

```bash
# 运行单个文件
go run main.go

# 运行多个文件
go run main.go utils.go

# 运行整个包
go run .

# 编译当前目录
go build

# 编译并指定输出文件名
go build -o app

# 编译指定包
go build github.com/user/project

# 编译多个包
go build github.com/user/project1  github.com/user/project2
```

## 测试

```bash
# 运行当前目录测试
go test

# 运行所有测试（递归）
go test ./...

# 显示详细输出
go test -v

# 显示覆盖率
go test -cover

# 显示基准
go test -bench=. -benchmem
```

# DataType

## 变量 (Variables)

### 变量定义与特点

- 变量是程序运行期间值可以改变的标识符
- 必须先声明后使用
- 具有特定的数据类型
- 存储在内存中，可以被读取和修改

### 变量声明方式

#### 1. 使用 `var` 关键字声明

```go
// 声明单个变量
var name string
var age int
var isActive bool

// 声明多个同类型变量
var x, y, z int

// 声明多个不同类型变量
var (
    name     string
    age      int
    isActive bool
)
```

#### 2. 声明并初始化

```go
// 指定类型并初始化
var name string = "张三"
var age int = 25

// 类型推导（省略类型）
var name = "张三"  // 自动推导为 string
var age = 25      // 自动推导为 int

// 多变量同时声明并初始化
var name, age = "张三", 25
```

#### 3. 短变量声明 `:=`

```go
// 只能在函数内部使用
name := "张三"
age := 25
isActive := true

// 多变量短声明
name, age := "张三", 25
```

### 变量零值

Go 语言中，已声明但未初始化的变量会被赋予零值：

```go
var i int      // 0
var f float64  // 0.0
var s string   // ""
var b bool     // false
var p *int     // nil
```

## 常量 (Constants)

### 常量定义与特点

- 常量是程序编译期间确定的值，运行期间不能改变
- 使用 `const` 关键字声明
- 必须在声明时进行初始化
- 只能是基本类型：布尔型、数字型、字符串型

### 常量声明方式

#### 1. 单个常量声明

```go
const PI = 3.14159
const AppName = "我的应用"
const MaxSize = 1000
```

#### 2. 多个常量声明

```go
const (
    StatusOK     = 200
    StatusError  = 500
    StatusNotFound = 404
)

// 类型推导
const (
    a = 1
    b = "hello"
    c = true
)
```

#### 3. 批量声明与 `iota`

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

// iota 的高级用法
const (
    _  = iota                    // 0，跳过
    KB = 1 << (10 * iota)       // 1 << 10 = 1024
    MB                          // 1 << 20 = 1048576
    GB                          // 1 << 30 = 1073741824
)
```

## 变量 vs 常量对比

| 特性       | 变量                   | 常量             |
| ---------- | ---------------------- | ---------------- |
| 关键字     | `var` 或 `:=`          | `const`          |
| 值是否可变 | 可变                   | 不可变           |
| 初始化时机 | 声明时可选，使用前必须 | 声明时必须       |
| 零值       | 有零值                 | 无零值概念       |
| 作用域     | 包级别或函数级别       | 包级别或函数级别 |
| 内存分配   | 运行时分配             | 编译时确定       |
| 支持类型   | 所有类型               | 基本类型         |

## 实际使用示例

```go
package main

import "fmt"

// 包级别常量
const (
    AppVersion = "1.0.0"
    MaxUsers   = 1000
)

// 包级别变量
var globalCounter int

func main() {
    // 函数内变量
    var userName string = "用户A"
    userAge := 30

    // 函数内常量
    const welcomeMsg = "欢迎使用我们的应用！"

    fmt.Printf("应用版本: %s\n", AppVersion)
    fmt.Printf("最大用户数: %d\n", MaxUsers)
    fmt.Printf("用户名: %s, 年龄: %d\n", userName, userAge)
    fmt.Println(welcomeMsg)

    // 变量可以修改
    userName = "用户B"
    userAge = 25
    globalCounter++

    fmt.Printf("修改后 - 用户名: %s, 年龄: %d\n", userName, userAge)
    fmt.Printf("全局计数器: %d\n", globalCounter)

    // 常量不能修改，以下代码会编译错误
    // AppVersion = "2.0.0"  // 编译错误
    // welcomeMsg = "再见"    // 编译错误
}
```

## 使用最佳实践

### 何时使用变量

1. 需要存储和修改的数据
2. 函数参数和返回值
3. 循环计数器
4. 临时计算结果

### 何时使用常量

1. 配置参数（如 API 密钥、版本号）
2. 数学常数（如 π、e）
3. 状态码、错误码
4. 固定的字符串消息
5. 枚举值

### 命名规范

```go
// 常量：大写字母开头，驼峰命名
const MaxRetryCount = 3
const DatabaseURL = "localhost:5432"

// 变量：小写字母开头，驼峰命名
var userName string
var isConnected bool

// 导出的变量和常量：大写字母开头
var GlobalConfig map[string]string
const PublicAPIVersion = "v2"
```

### 性能考虑

- 常量在编译时确定，性能更好
- 对于不会改变的值，优先使用常量
- 常量不占用运行时内存

## 常见错误

```go
// 错误示例
const dynamicValue = time.Now().Unix() // 编译错误：不能使用动态值
var uninitializedConst const           // 语法错误：常量必须初始化

// 正确做法
var dynamicValue = time.Now().Unix()   // 变量可以使用动态值
const staticValue = 123456789          // 常量使用编译时确定的值
```

## 总结

变量和常量是 Go 语言的基础概念：

- **变量**适用于需要改变的数据，提供灵活性
- **常量**适用于固定不变的值，提供性能和安全性
- 合理使用两者可以让代码更清晰、更高效、更安全

## Go 语言基础数据类型与方法详解

### 1. 数字（int&float）

#### 1.1 整数类型

Go 语言提供了多种整数类型，分为有符号和无符号两大类：

```go
// 有符号整数
var i8 int8 = -128     // 8位，范围：-128 到 127
var i16 int16 = -32768 // 16位，范围：-32768 到 32767
var i32 int32 = -2147483648 // 32位
var i64 int64 = -9223372036854775808 // 64位

// 无符号整数
var u8 uint8 = 255     // 8位，范围：0 到 255
var u16 uint16 = 65535 // 16位，范围：0 到 65535
var u32 uint32 = 4294967295 // 32位
var u64 uint64 = 18446744073709551615 // 64位

// 平台相关类型
var i int = 42    // 32位或64位，取决于平台
var u uint = 42   // 32位或64位，取决于平台
var ptr uintptr   // 存储指针的无符号整数

// 类型别名
var b byte = 255    // uint8的别名
```

#### 1.2 浮点数类型

```go
var f32 float32 = 3.14159
var f64 float64 = 3.141592653589793

// 复数类型
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

#### 1.3 数值类型常用方法和操作

##### 类型转换

```go
package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {
    // 类型转换
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(i)

    // 字符串转数值
    str := "123"
    num, err := strconv.Atoi(str)           // 字符串转int
    if err == nil {
        fmt.Println("Converted:", num)
    }

    f64, err := strconv.ParseFloat("3.14", 64) // 字符串转float64
    i64, err := strconv.ParseInt("123", 10, 64) // 字符串转int64
}
```

##### 数学运算

```go
import "math"

// 常用数学函数
abs := math.Abs(-3.14)      // 绝对值
ceil := math.Ceil(3.14)     // 向上取整
floor := math.Floor(3.14)   // 向下取整
round := math.Round(3.14)   // 四舍五入
sqrt := math.Sqrt(16)       // 平方根
pow := math.Pow(2, 3)       // 幂运算
sin := math.Sin(math.Pi/2)  // 三角函数

// 最大最小值
max := math.Max(1.0, 2.0)
min := math.Min(1.0, 2.0)
```

### 2. 布尔（bool）

布尔类型只有两个值：`true` 和 `false`

```go
var b1 bool = true
var b2 bool = false
var b3 bool          // 零值为false

// 布尔运算
result1 := true && false    // 逻辑与
result2 := true || false    // 逻辑或
result3 := !true           // 逻辑非

// 比较运算返回布尔值
equal := (5 == 5)          // true
notEqual := (5 != 3)       // true
greater := (5 > 3)         // true
less := (3 < 5)            // true
greaterEqual := (5 >= 5)   // true
lessEqual := (3 <= 5)      // true

// 布尔值与字符串转换
import "strconv"

boolVal, err := strconv.ParseBool("true") // true, nil

// 布尔值在条件语句中的使用
if b1 {
    fmt.Println("b1 is true")
}

// 布尔值不能与数字直接转换（与C不同）
// var i int = int(true)  // 编译错误
```

### 3. 字符串（string）

字符串是不可变的字节序列。

#### 4.1 字符串基础操作

```go
// 字符串声明
var s1 string = "Hello"
var s2 string = `多行
字符串
使用反引号`

// 字符串长度（字节数）
length := len(s1)

// 字符串索引（获取字节）
b := s1[0]  // 'H'的字节值

// 字符串切片
sub := s1[1:4]  // "ell"
```

#### 4.2 strings 包常用方法

```go
import "strings"

s := "Hello, World! Hello!"

s[0]
s[1:3]

// 查找
contains := strings.Contains(s, "World")        // true index := strings.Index(s, "World")  lastIndex := strings.LastIndex(s, "Hello")
hasPrefix := strings.HasPrefix(s, "Hello")      // true
hasSuffix := strings.HasSuffix(s, "!")          // true

count := strings.Count(s, "l")                  // 3

// 大小写
upper := strings.ToUpper(s)                     // "HELLO, WORLD! HELLO!"
lower := strings.ToLower(s)                     // "hello, world! hello!"
title := strings.Title("hello world")          // "Hello World"

replaced := strings.Replace(s, "Hello", "Hi", 1)  // "Hi, World! Hello!"
replacedAll := strings.ReplaceAll(s, "Hello", "Hi") // "Hi, World! Hi!"

// 修剪
trimPrefix := strings.TrimPrefix(s, "Hello")   // ", World! Hello!"
trimSuffix := strings.TrimSuffix(s, "!")       // "Hello, World! Hello"
trimmed := strings.TrimSpace("  hello  ")      // "hello"

// 连接/分割
joined := strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"

parts := strings.Split("a,b,c", ",")           // ["a", "b", "c"]
fields := strings.Fields("  hello   world  ")  // ["hello", "world"]
```

#### 4.3 字符串格式化

```go
import "fmt"

name := "Alice"
age := 30
height := 5.6

// 格式化动词
formatted := fmt.Sprintf("Name: %s, Age: %d, Height: %.1f", name, age, height)
fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

// 常用格式化动词
fmt.Printf("%v\n", 42)          // 默认格式
fmt.Printf("%+v\n", struct{X int}{42})
fmt.Printf("%#v\n", []int{1,2})

fmt.Printf("%T\n", 42)          // 类型
fmt.Printf("%%\n")              // 百分号字面量

// 数字格式化
fmt.Printf("%d\n", 42)          // 十进制
fmt.Printf("%b\n", 42)          // 二进制
fmt.Printf("%o\n", 42)          // 八进制
fmt.Printf("%x\n", 42)          // 十六进制
fmt.Printf("%X\n", 42)          // 大写十六进制
fmt.Printf("%f\n", 3.14159)     // 浮点数
fmt.Printf("%.2f\n", 3.14159)   // 保留2位小数

// 字符串格式化
fmt.Printf("%s\n", "hello")           // 按原样输出字符串: hello
fmt.Printf("%q\n", "hello")           // 带双引号的字符串: "hello"

char := '世'
fmt.Printf("%c\n", char)              // 以字符方式输出: 世，会将整数（如rune、byte）按照字符输出
fmt.Printf("%U\n", char)              // Unicode格式: U+4E16


fmt.Printf("%8s\n", "hi")             // 宽度为8，右对齐: "      hi"
fmt.Printf("%-8s*\n", "hi")           // 左对齐，宽度8: "hi      *"
fmt.Printf("%.3s\n", "abcdefg")       // 最多输出3个字符: abc

// 错误
baseErr := errors.New("database connection failed")
wrappedErr := fmt.Errorf("failed to fetch user: %w", baseErr) // 保留错误链

errors.Unwrap(wrappedErr)
```

### 4. 正则（regexp）

Go 的正则表达式包提供了强大的模式匹配功能。

#### 5.1 基础用法

```go
import "regexp"

// 编译正则表达式
pattern := `\d+`  // 匹配一个或多个数字
re, err := regexp.Compile(pattern)
if err != nil {
    log.Fatal(err)
}

// 或者使用MustCompile（如果模式错误会panic）
re := regexp.MustCompile(`\d+`)

text := "I have 123 apples and 456 oranges"

// 是否匹配
found := re.MatchString(text)               // true

// 查找匹配
first := re.FindString(text)                // "123" re.FindStringIndex(text)
all := re.FindAllString(text, -1)           // ["123", "456"] re.FindAllStringIndex(text, -1)
```

#### 5.2 子匹配和分组

```go
// 使用分组
re := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)  // 匹配邮箱

text := "Contact us: john@example.com or alice@test.org"

// 查找子匹配 []string john@example.com  john example.com
submatch := re.FindStringSubmatch(text)

// 所有匹配及其分组 [][]string
allSubmatch := re.FindAllStringSubmatch(text, -1)

// 命名分组
re := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+\.\w+)`)

names := re.SubexpNames()
matches := re.FindAllStringSubmatch(text, -1)

for _, match := range matches {
    for i, name := range names {
        if i != 0 && name != "" {
            fmt.Printf("%s: %s\n", name, match[i])
        }
    }
    fmt.Println("---")
}
```

#### 5.3 替换/分割

```go
re := regexp.MustCompile(`(\d+)`)
text := "I have 123 apples and 456 oranges"

// 简单替换
replaced := re.ReplaceAllString(text, "XXX")

// 分组替换
masked := re.ReplaceAllString(emailText, "$1@***.**")

// 函数替换
replacedFunc := re.ReplaceAllStringFunc(text, func(s string) string {
    num, _ := strconv.Atoi(s)

    return strconv.Itoa(num * 2)
})


// 基本分割
re := regexp.MustCompile(`[,\s]+`)
text := "Go,Python Java,Rust"

fields := re.Split(text, -1)  // []string: ["Go", "Python", "Java", "Rust"]
```

### 5. 数组 (arrays)

数组是固定长度的同类型元素序列。

#### 声明和初始化

```go
// 声明数组
var arr1 [5]int                    // 零值初始化
var arr2 = [5]int{1, 2, 3, 4, 5}   // 字面量初始化
arr3 := [5]int{1, 2, 3, 4, 5}      // 简短声明
arr4 := [...]int{1, 2, 3, 4, 5}    // 自动推断长度
arr5 := [5]int{1: 10, 3: 30}       // 指定索引初始化
```

#### 数组操作

```go
func arrayOperations() {
    arr := [5]int{1, 2, 3, 4, 5}

    // 访问元素
    fmt.Println(arr[0])        // 1
    arr[0] = 10                // 修改元素

    // 获取长度
    fmt.Println(len(arr))      // 5

    // 遍历数组
    for i, v := range arr {
        fmt.Printf("索引: %d, 值: %d\n", i, v)
    }

    // 数组比较（相同类型 + 长度的数组可以比较）
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    fmt.Println(arr1 == arr2)  // true
}
```

### 6. 切片 (slices)

切片是动态数组，是对数组的抽象。

#### 声明和初始化

```go
// 声明切片
var slice1 []int               // nil切片
slice2 := []int{1, 2, 3, 4, 5} // 字面量初始化
slice3 := make([]int, 5)       // 使用make创建，长度为5
slice4 := make([]int, 3, 5)    // 长度为3，容量为5

// 从数组创建切片
arr := [5]int{1, 2, 3, 4, 5}

slice5 := arr[1:4]             // [2, 3, 4]
```

#### 切片操作

```go
func sliceOperations() {
    slice := []int{1, 2, 3, 4, 5}

    // 基本信息
    fmt.Printf("长度: %d, 容量: %d\n", len(slice), cap(slice))

    // 追加元素
    slice = append(slice, 6, 7, 8)
    slice = append(slice, []int{9, 10}...) // 追加另一个切片

    // 复制切片
    newSlice := make([]int, len(slice))
    copy(newSlice, slice)

    // 插入元素（在索引2处插入99）
    index := 2
    slice = append(slice[:index], append([]int{99}, slice[index:]...)...)

    // 删除元素（删除索引2的元素）
    slice = append(slice[:index], slice[index+1:]...)

    // 切片切片
    subSlice := slice[1:4]

    // 遍历切片
    for i, v := range slice {
        fmt.Printf("索引: %d, 值: %d\n", i, v)
    }
}
```

#### 切片的内存布局

```go
func sliceInternals() {
    slice := make([]int, 3, 5)
    fmt.Printf("地址: %p, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))
}
```

### 7. 映射 (maps)

映射是键值对的无序集合。

#### 声明和初始化

```go
// 声明映射
var map1 map[string]int                    // nil映射
map2 := make(map[string]int)               // 空映射
map3 := map[string]int{                    // 字面量初始化
    "apple":  5,
    "banana": 3,
    "orange": 8,
}
map4 := make(map[string]int, 10)           // 指定初始容量
```

#### 映射操作

```go
func mapOperations() {
    m := make(map[string]int)

    // 添加/修改元素
    m["apple"] = 5
    m["banana"] = 3
    m["orange"] = 8

    // 访问元素
    value := m["apple"]                    // 获取值
    value, ok := m["apple"]                // 检查键是否存在
    if ok {
        fmt.Printf("apple: %d\n", value)
    }

    // 删除元素
    delete(m, "banana")

    // 获取长度
    fmt.Printf("映射长度: %d\n", len(m))

    // 遍历映射
    for key, value := range m {
        fmt.Printf("键: %s, 值: %d\n", key, value)
    }

    // 只遍历键
    for key := range m {
        fmt.Printf("键: %s\n", key)
    }

    // 只遍历值
    for _, value := range m {
        fmt.Printf("值: %d\n", value)
    }

    // set
    set := make(map[string]struct{}) // 不存储任何数据 var s struct{} unsafe.Sizeof(s) == 0  存储任意类型的值 var i interface{} unsafe.Sizeof(i) == 16
    set["apple"] = struct{}{}
    set["banana"] = struct{}{}

    if _, ok := set["apple"]; ok { // 检查元素是否存在
        fmt.Println("apple exists")
    }

    delete(set, "apple") // 删除元素
}
```

### 8. 结构体 (structs)

#### 定义和初始化

```go
type UserMap map[int]string // 类型定义，可以为新类型添加方法
type UserMap = map[int]string // 类型别名，不可以为新类型添加方法

type Number interface { // 联合类型
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

type Pair[K, V any] struct { //  泛型
    Key   K
    Value V
}

// 定义结构体
type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address  // 嵌套结构体
}

type Address struct {
    Street   string
    City     string
    ZipCode  string
}

// 初始化结构体
func structInitialization() {
    // 零值初始化
    var p1 Person

    // 字面量初始化
    p2 := Person{
        Name:  "张三",
        Age:   30,
        Email: "zhangsan@example.com",
    }

    // 位置参数初始化
    p3 := Person{"李四", 25, "lisi@example.com", Address{}}

    // 使用new
    p4 := new(Person)
    p4.Name = "王五"
}
```

#### 结构体方法

```go
// 值接收者方法
func (p Person) GetInfo() string {
    return fmt.Sprintf("姓名: %s, 年龄: %d", p.Name, p.Age)
}

// 指针接收者方法
func (p *Person) SetAge(age int) {
    p.Age = age
}

func (p *Person) HaveBirthday() {
    p.Age++
}

// 使用方法
func usePersonMethods() {
    p := Person{Name: "张三", Age: 30}

    // 调用值接收者方法
    info := p.GetInfo()
    fmt.Println(info)

    // 调用指针接收者方法
    p.SetAge(31)
    p.HaveBirthday()

    fmt.Printf("新年龄: %d\n", p.Age)
}
```

#### json

```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id" db:"user_id"`
    Name     string `json:"name" db:"user_name"`
    Email    string `json:"email,omitempty" db:"email"` // 空值时忽略
    Password string `json:"-" db:"password"`  // 序列化时忽略
}

func structTags() {
    user := User{
        ID:       1,
        Name:     "张三",
        Email:    "zhangsan@example.com",
        Password: "secret",
    }

    // JSON序列化 []byte
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))

    // JSON反序列化 []byte
    var newUser User
    json.Unmarshal(jsonData, &newUser)
    fmt.Printf("%+v\n", newUser)
}
```

#### 结构体嵌入

```go
// 匿名字段（嵌入）
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() {
    fmt.Printf("%s 发出声音\n", a.Name)
}

type Dog struct {
    Animal        // 嵌入Animal
    Breed string
}

func (d Dog) Bark() {
    fmt.Printf("%s 汪汪叫\n", d.Name)
}

// 接口实现
type Speaker interface {
    Speak()
}

func structComposition() {
    dog := Dog{
        Animal: Animal{Name: "旺财", Age: 3},
        Breed:  "金毛",
    }

    // 可以直接访问嵌入结构体的字段和方法
    fmt.Println(dog.Name)  // 直接访问Animal的Name字段
    dog.Speak()           // 调用Animal的方法
    dog.Bark()            // 调用Dog自己的方法

    // 接口使用
    var speaker Speaker = dog
    speaker.Speak()
}
```

#### 接口（Interface）

```go
// 结构体是具体类型
type Person struct {
    Name  string
    Age   int
    Email string
}

// 接口是抽象类型
type Greeter interface {
    Greet() string
}

func (p Person) Greet() string {
    return "你好，我是" + p.Name
}

p := Person{
    Name:  "张三",
    Age:   30,
    Email: "zhangsan@example.com",
}

p.Greet()
```

### 9. 指针 (pointers)

#### 指针基础

```go
func pointerBasics() {
    var x int = 42
    var p *int        // 声明指向int的指针

    p = &x           // 获取x的地址
    fmt.Printf("x的值: %d\n", x)
    fmt.Printf("x的地址: %p\n", &x)
    fmt.Printf("p的值: %p\n", p)
    fmt.Printf("p指向的值: %d\n", *p)  // 解引用

    // 通过指针修改值
    *p = 100
    fmt.Printf("修改后x的值: %d\n", x)

    // 零值指针
    var ptr *int
    if ptr == nil {
        fmt.Println("ptr是nil指针")
    }

    // 使用new分配内存
    ptr = new(int)
    *ptr = 50
    fmt.Printf("new分配的值: %d\n", *ptr)

    slice := make([]int, 5)     // 返回 []int
    m := make(map[string]int)   // 返回 map[string]int
    ch := make(chan int)        // 返回 chan int

    // new 创建指针
    ptr := new(int)             // 返回 *int
    slicePtr := new([]int)      // 返回 *[]int
    mapPtr := new(map[string]int) // 返回 *map[string]int
}
```

#### 指针与函数

```go
// 值传递
func modifyValue(x int) {
    x = 100
}

// 指针传递
func modifyPointer(x *int) {
    *x = 100
}

// 返回指针
func createInt(value int) *int {
    x := value
    return &x  // 返回局部变量的地址是安全的
}

func pointerAndFunction() {
    x := 42

    modifyValue(x)
    fmt.Printf("值传递后: %d\n", x)    // 42，未改变

    modifyPointer(&x)
    fmt.Printf("指针传递后: %d\n", x)  // 100，已改变

    ptr := createInt(200)
    fmt.Printf("创建的指针值: %d\n", *ptr)
}
```

### 10. 时间和日期处理 (time)

#### 1. time 包基础

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建特定时间
    specificTime := time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC)
    fmt.Println("指定时间:", specificTime)

    // 获取当前时间
    now := time.Now()
    fmt.Println("当前时间:", now)

    // 解析字符串为时间
    stringTime := "2024-01-15 14:30:00"
    parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
    if err != nil {
        fmt.Println("解析错误:", err)
        return
    }

    fmt.Println("自定义:", now.Format("2006-01-02 15:04:05"))
    fmt.Println("日期:", now.Format("2006年01月02日"))
    fmt.Println("时间:", now.Format("15:04:05"))

    // 时间戳 sec nsec
    timestampTime := time.Unix(1705300200, 0)

    timestamp := now.Unix() // int64
}
```

#### 3. 时间计算和比较

```go
func timeCalculations() {
    now := time.Now()

    // 时间加减
    future := now.Add(2 * time.Hour)
    past := now.Add(-1 * time.Hour)

    fmt.Println("现在:", now.Format("15:04:05"))
    fmt.Println("2小时后:", future.Format("15:04:05"))
    fmt.Println("1小时前:", past.Format("15:04:05"))

    // 时间差
    duration := time2.Sub(time1) // time.Since(time1) == time.Now().Sub(time1)
    fmt.Println("时间差:", duration)
    fmt.Println("小时差:", duration.Hours()) // float64
    fmt.Println("分差:", duration.Minutes()) // float64
    fmt.Println("秒差:", duration.Seconds()) // float64

    fmt.Println("差:", duration.Milliseconds()) // int64
    fmt.Println("差:", duration.Microseconds()) // int64
    fmt.Println("差:", duration.Nanoseconds()) // int64

    duration := time.ParseDuration("1h30m")

    // 时间比较
    time1 := time.Now()
    time2 := time1.Add(1 * time.Hour)

    fmt.Println("time1 < time2:", time1.Before(time2))
    fmt.Println("time1 > time2:", time1.After(time2))
    fmt.Println("time1 == time2:", time1.Equal(time2))

    // 日期计算
    nextWeek := now.AddDate(0, 0, 7)  // 年, 月, 日
    nextMonth := now.AddDate(0, 1, 0)
    nextYear := now.AddDate(1, 0, 0)

    fmt.Println("下周:", nextWeek.Format("2006-01-02"))
    fmt.Println("下月:", nextMonth.Format("2006-01-02"))
    fmt.Println("明年:", nextYear.Format("2006-01-02"))
}
```

### 11. 错误处理 (errors)

#### 1. 基本错误处理

```go
import (
    "errors"
    "fmt"
)

// 创建错误
func basicErrorHandling() {
    // 使用errors.New创建错误
    err1 := errors.New("这是一个简单错误")
    fmt.Println("错误1:", err1)

    // 使用fmt.Errorf创建格式化错误
    name := "用户名"
    err2 := fmt.Errorf("用户 %s 不存在", name)
    fmt.Println("错误2:", err2)

    // 函数返回错误
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("除法错误:", err)
    } else {
        fmt.Println("结果:", result)
    }
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}
```

#### 2. 自定义错误类型

```go
// 自定义错误类型
type ValidationError struct {
    Field   string
    Message string
    Code    int
}

// 实现Error接口
func (e *ValidationError) Error() string {
    return fmt.Sprintf("验证错误 [%s]: %s (代码: %d)", e.Field, e.Message, e.Code)
}

func validateUser(name string, age int) error {
    if name == "" {
        return &ValidationError{
            Field:   "name",
            Message: "用户名不能为空",
            Code:    1001,
        }
    }
    if age < 18 {
        return &ValidationError{
            Field:   "age",
            Message: "年龄必须大于18岁",
            Code:    1002,
        }
    }
    return nil
}
```

#### 3. 错误包装和链式处理

```go
import (
    "fmt"
)

// Go 1.13+ 的错误包装
func errorWrapping() {
    err := processData()
    if err != nil {
        fmt.Println("处理失败:", err)

        // 检查错误链中是否包含特定错误
        if errors.Is(err, sql.ErrNoRows) {
            fmt.Println("没有找到数据")
        }

        // 获取错误链中的特定类型
        var validationErr *ValidationError
        if errors.As(err, &validationErr) {
            fmt.Printf("发现验证错误: %s\n", validationErr.Field)
        }
    }
}

func processData() error {
    err := fetchDataFromDB()
    if err != nil {
        return fmt.Errorf("处理数据时出错: %w", err)
    }
    return nil
}

func fetchDataFromDB() error {
    err := connectDB()
    if err != nil {
        return fmt.Errorf("获取数据失败: %w", err)
    }
    return nil
}

func connectDB() error {
    return &ValidationError{
        Field:   "connection",
        Message: "数据库连接失败",
        Code:    2001,
    }
}
```

### 12. 枚举 (enum)

Go 语言虽然没有内置的 enum 类型，但提供了多种方式来实现枚举功能。以下是 Go 语言中实现枚举的详细介绍：

#### 1. 使用 iota 实现整数枚举

##### 基本用法

```go
package main

import "fmt"

// 数字
type Color int

const (
    Red Color = 0  // 0
    Green Color = 1       // 1
    Blue Color = 2         // 2
)

// 字符串
func main() {
    fmt.Println(Red, Green, Blue) // 输出: 0 1 2
}

type Status string

const (
    StatusPending   Status = "pending"
    StatusRunning   Status = "running"
    StatusCompleted Status = "completed"
    StatusFailed    Status = "failed"
)
```

### 13. 类型 (.())

#### 断言 (Type Assertion)

```go
var i interface{} = "hello"

// 类型断言
s := i.(string)
fmt.Println(s) // hello

// 安全类型断言（推荐）
if s, ok := i.(string); ok {
    fmt.Println("是字符串:", s)
} else {
    fmt.Println("不是字符串")
}

// 类型断言
func checkType(i interface{}) {
    switch v, t := i.(type); t {
    case int:
        fmt.Printf("整数: %d\n", v)
    case string:
        fmt.Printf("字符串: %s\n", v)
    case bool:
        fmt.Printf("布尔值: %t\n", v)
    case []int:
        fmt.Printf("整数切片: %v\n", v)
    default:
        fmt.Printf("未知类型: %T\n", v)
    }
}

import (
    "fmt"
    "reflect"
)

func checkType(i interface{}) {
    v := reflect.ValueOf(val)
    t := v.Type()

    switch t.Kind() {
    case reflect.Int:
        fmt.Printf("是 int: %d\n", v.Int())
    case reflect.String:
        fmt.Printf("是 string: %s\n", v.String())
    case reflect.Bool:
        fmt.Printf("是 bool: %v\n", v.Bool())
    case reflect.Slice:
        fmt.Printf("是切片, 长度 %d\n", v.Len())
    case reflect.Struct:
        fmt.Println("是结构体, 字段数:", v.NumField())
    default:
        fmt.Println("其他类型:", t.Kind())
    }
}

// 接口断言
type Writer interface {
    Write([]byte) (int, error)
}

func checkInterface(i interface{}) {
    if _, ok := i.(Writer); ok {
        fmt.Println("实现了Writer接口")
    } else {
        fmt.Println("没有实现Writer接口")
    }
}
```

## Go 语言条件语句与循环详解

### 1. 条件语句

#### 1. if 语句

```go
if condition1 {
    // condition1 为真时执行
} else if condition2 {
    // condition2 为真时执行
} else {
    // 所有条件都为假时执行
}

if v := math.Pow(x, n); v < lim {
    return v
}
// v 的作用域仅在 if-else 块内
```

#### 2. switch 语句

```go
switch expression {
case value1:
    // 执行语句
case value2:
    // 执行语句
default:
    // 默认执行语句
}

switch day {
case 1:
    fmt.Println("星期一")
case 2:
    fmt.Println("星期二")
case 3, 4, 5:
    fmt.Println("工作日")
case 6, 7:
    fmt.Println("周末")
default:
    fmt.Println("无效的日期")
}

switch num := rand.Intn(10); num {
case 0, 2, 4, 6, 8:
    fmt.Printf("%d 是偶数\n", num)
case 1, 3, 5, 7, 9:
    fmt.Printf("%d 是奇数\n", num)
}

switch {
case score >= 90:
    fmt.Println("优秀")
case score >= 80:
    fmt.Println("良好")
case score >= 60:
    fmt.Println("及格")
default:
    fmt.Println("不及格")
}
```

### 2. 循环语句

> Go 语言只有一种循环结构：`for` 循环，但它有多种形式。

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

for i := range 10 {
    fmt.Println(i)
}

for range 10 {}

str := "Hello世界"
for i := range len(str) {
    fmt.Printf("%c ", str[i])  // 遍历字节，可能乱码
}
for i, v := range str {
    fmt.Printf("位置 %d: %c\n", i, v) // 遍历 Unicode 字符
}

arr := []int{1, 2, 3, 4, 5}
for i, v := range arr {
    fmt.Printf("索引: %d, 值: %d\n", i, v)
}
for _, value := range arr {
    fmt.Println(value)
}
for i := range arr {
    fmt.Println(i)
}

m := map[string]int{
    "apple":  5,
    "banana": 3,
    "orange": 2,
}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
for key := range m {
    fmt.Println(key)
}

ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)
for v := range ch {
    fmt.Println(v)
}

// while
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}

// while + loop
for {
    fmt.Println("无限循环")

    // 需要 break 或 return 来退出
    if someCondition {
        break
    }
}
```

## Go 语言函数详解

### 1. 基本函数语法

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // 函数体
    return value
}

func add(a int, b int) int {
    return a + b
}

// 简化写法（相同类型的参数）
func add(a, b int) int {
    return a + b
}

// 多返回值函数
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println(quotient, remainder) // 3 1
}

// 返回值命名
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)

    return area, perimeter
    return // 自动返回命名的返回值
}

// 可变参数函数
values := []int{1, 2, 3, 4}
sum(values...)

func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 函数作为值
calculate(10, 5, multiply) // 50

type Op func(int, int) int

func calculate(a, b int, op Op) int {
    return op(a, b)
}

func multiply(a, b int) int {
    return a * b
}

// 闭包
multiplier := createMultiplier(3)
multiplier(4) // 12

func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// 高阶函数
func Map[T, R any](slice []T, fn func(T) R) []R {
    result := make([]R, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}
squares := Map(numbers, func(n int) int {
        return n * n
})

func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := []T{}
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
evens := Filter([]int{1, 2, 3, 4}, func(n int) bool {
    return n%2 == 0
})

func Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

sum := Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})

// 结构体
type Rectangle struct {
    Width  float64
    Height float64
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
rect.Scale(2)

// defer 延迟函数的执行，stack
func main() {
    defer fmt.Println("这会最后执行")
    defer fmt.Println("这会倒数第二执行")

    fmt.Println("这会首先执行")
}

// 先检查错误，确认文件成功打开后才执行 defer
// 如果打开失败，file 是 nil，不会调用 Close()
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}
defer file.Close()



// 错误处理
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
```

# I/O

## 文件读取

```go
// os.ReadFile 读取全部内容（简单)
content, err := os.ReadFile("example.txt")
if err != nil {
    fmt.Println("读取失败:", err)
    return
}
fmt.Println(string(content))

// io.ReadAll 适合更大自定义控制
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}
defer file.Close()

content, err = io.ReadAll(file)
if err != nil {
    fmt.Println("读取失败:", err)
    return
}
fmt.Println(string(content))

// bufio.Scanner 行读取
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println("每行:", scanner.Text())
}

if err := scanner.Err(); err != nil {
    fmt.Println("读取文件失败:", err)
    return
}

// bufio.Reader 按分隔符读取
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}

reader := bufio.NewReader(file)
for {
    content, err := reader.ReadString('\n')
    if err != nil {
        if err == io.EOF {
            if len(content) > 0 {
                fmt.Print(content)
            }
            break
        }
        fmt.Println("读取错误:", err)
        return
    }
    fmt.Print(content)
}

// file.Read 按块读取
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}

buffer := make([]byte, 1024)
for {
    _, err := file.Read(buffer)
    if err != nil {
        if err == io.EOF {
            break
        }
        fmt.Println("读取错误:", err)
        return
    }
    fmt.Print(string(buffer))
}
```

## 文件写入

```go
// os.WriteFile 整体写入
content := []byte("写入内容")
err := os.WriteFile("out.txt", content, 0o644)
if err != nil {
    fmt.Println("写入失败:", err)
    return
}
fmt.Println("写入成功")

// io.WriteString / io.Write example
file, err := os.Create("io_output.txt")
if err != nil {
    fmt.Println("创建文件失败:", err)
    return
}
defer file.Close()

_, err = io.WriteString(file, "io 写入的字符串内容\n")
if err != nil {
    fmt.Println("写入失败:", err)
    return
}

_, err = io.Write(file, []byte("io 写入的字符串内容\n"))
if err != nil {
    fmt.Println("写入失败:", err)
    return
}

// bufio.NewWriter 缓冲写入
file, err := os.Create("buffered.txt")
if err != nil {
    fmt.Println("创建文件失败:", err)
    return
}
defer file.Close()

writer := bufio.NewWriter(file)
for i := 0; i < 10; i++ {
    fmt.Fprintf(writer, "第 %d 行\n", i+1)
}
err = writer.Flush()
if err != nil {
    fmt.Println("刷新缓冲区失败:", err)
    return
}
fmt.Println("缓冲写入完成")

// file.WriteString / file.Write
file, err := os.Create("output.txt")
if err != nil {
    fmt.Println("创建文件失败:", err)
    return
}
defer file.Close()

_, err = file.WriteString("字符串内容")
if err != nil {
    fmt.Println("写入失败:", err)
    return
}
_, err = file.Write([]byte("二进制内容"))
if err != nil {
    fmt.Println("写入失败:", err)
    return
}

// os.OpenFile 追加写入
file, err = os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o644)
if err != nil {
    fmt.Println("打开文件失败:", err)
    return
}
defer file.Close()
```

## 删除/重命名/复制

```go
func file() {
    // 删除文件
    err := os.Remove("unwanted.txt")
    if err != nil {
        fmt.Println("删除文件失败:", err)
    }

    // 递归删除目录及其内容
    err = os.RemoveAll("path/to/dir")
    if err != nil {
        fmt.Println("递归删除失败:", err)
    }

    // 重命名
    err := os.Rename("oldname.txt", "newname.txt")
    if err != nil {
        fmt.Println("重命名失败:", err)
        return
    }

    // 复制
    func copyFile(src, dst string) error {
        sourceFile, err := os.Open(src)
        if err != nil {
            return err
        }
        defer sourceFile.Close()

        destFile, err := os.Create(dst)
        if err != nil {
            return err
        }
        defer destFile.Close()

        _, err = io.Copy(destFile, sourceFile)
        return err
    }

    func copyFileExample() {
        err := copyFile("source.txt", "destination.txt")
        if err != nil {
            fmt.Println("复制文件失败:", err)
            return
        }
    }
}
```

## 文件信息

```go
// 判断 error 是否为 os.ErrNotExist
func isFileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

func getFileStat() {
    stat, err := os.Stat("example.txt")
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("文件不存在")
        } else {
            fmt.Println("获取文件信息失败:", err)
        }
        return
    }

    fmt.Println("文件名:", stat.Name())
    fmt.Println("文件大小:", stat.Size(), "字节")
    fmt.Println("修改时间:", stat.ModTime())
    fmt.Println("是否为目录:", stat.IsDir())
    fmt.Println("文件权限:", stat.Mode() & 0o777)
}
```

## 文件夹

```go
func directory() {
    // 创建单级目录
    err := os.Mkdir("dir", 0o755)
    if err != nil && !os.IsExist(err) {
        fmt.Println("创建目录失败:", err)
        return
    }

    // 创建多级目录
    err = os.MkdirAll("path/to/nested/dir", 0o755)
    if err != nil {
        fmt.Println("创建多级目录失败:", err)
        return
    }

    // 读取目录
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("读取目录失败:", err)
        return
    }

    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Printf("[目录] %s\n", entry.Name())
        } else {
            fmt.Printf("[文件] %s\n", entry.Name())
        }
    }
}
```

# Network

## HTTP 服务端 (Server)

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // 最简单的处理器
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // 启动服务器
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### http.ResponseWriter

**ResponseWriter 是一个接口，用于构造 HTTP 响应：**

```go
type ResponseWriter interface {
    // 设置响应头
    Header() http.Header

    // 写入状态码
    WriteHeader(statusCode int)

    // 写入响应体（自动设置 200 状态码）
    Write([]byte) (int, error)
}
```

#### 基本用法

**1. 写入文本响应：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
```

**2. 设置状态码：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound) // 404
    w.Write([]byte("Page not found"))
}
```

**3. 设置响应头：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "value")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "success"}`))
}
```

**4. JSON 响应：**

```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "name": "John",
        "age":  30,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### \*http.Request

**Request 包含客户端请求的所有信息：**

#### 常用字段

```go
type Request struct {
    Method string              // GET, POST, PUT, DELETE, etc.
    URL    *url.URL           // 请求的 URL
    Proto  string             // "HTTP/1.1"
    Header Header             // 请求头
    Body   io.ReadCloser      // 请求体
    Host   string             // 主机名
    Form   url.Values         // 解析后的表单数据
    PostForm url.Values       // POST 表单数据
    RemoteAddr string         // 客户端地址
    Context context.Context   // 请求上下文
    // ... 还有更多字段
}
```

#### 读取请求信息

**1. 读取 URL 参数：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // URL: /search?q=golang&page=1

    query := r.URL.Query()
    q := query.Get("q")           // "golang"
    page := query.Get("page")     // "1"

    // 获取所有值（如果有多个相同键）
    tags := query["tag"]          // []string

    fmt.Fprintf(w, "Query: %s, Page: %s", q, page)
}
```

**2. 读取 URL 路径：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path           // "/users/123"
    fmt.Fprintf(w, "Path: %s", path)
}
```

**3. 读取请求头：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    contentType := r.Header.Get("Content-Type")
    userAgent := r.Header.Get("User-Agent")
    auth := r.Header.Get("Authorization")

    fmt.Fprintf(w, "Content-Type: %s", contentType)
}
```

**4. 读取 POST 表单数据：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 必须先调用 ParseForm
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    fmt.Fprintf(w, "Username: %s", username)
}
```

**5. 读取 JSON 请求体：**

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var user User

    // 解码 JSON
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    fmt.Fprintf(w, "User: %s, Email: %s", user.Name, user.Email)
}
```

**6. 读取原始请求体：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    fmt.Fprintf(w, "Body: %s", string(body))
}
```

**7. 检查请求方法：**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 或使用 switch
    switch r.Method {
    case http.MethodGet:
        // 处理 GET
    case http.MethodPost:
        // 处理 POST
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

### 完整示例

**1. RESTful API 示例：**

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        // GET /users - 获取所有用户
        json.NewEncoder(w).Encode(users)

    case http.MethodPost:
        // POST /users - 创建新用户
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        user.ID = len(users) + 1
        users = append(users, user)

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // 从 URL 提取 ID: /users/1
    path := strings.TrimPrefix(r.URL.Path, "/users/")
    id, err := strconv.Atoi(path)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // 查找用户
    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }

    http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)

    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**2. 文件上传示例：**

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 解析 multipart form (最大 10MB)
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 获取文件
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to get file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fmt.Fprintf(w, "Uploaded file: %s, Size: %d bytes",
        handler.Filename, handler.Size)
}
```

**3. Cookie 操作：**

```go
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:     "session_id",
        Value:    "abc123",
        Path:     "/",
        MaxAge:   3600, // 1小时
        HttpOnly: true,
        Secure:   true,
    }
    http.SetCookie(w, cookie)
    w.Write([]byte("Cookie set"))
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_id")
    if err != nil {
        http.Error(w, "Cookie not found", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Cookie value: %s", cookie.Value)
}
```

**4. 重定向：**

```go
func redirectHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/new-location", http.StatusFound) // 302
    // 或 http.StatusMovedPermanently (301)
}
```

**5. 使用 Context：**

```go
func timeoutHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(2 * time.Second):
        w.Write([]byte("Request completed"))
    case <-ctx.Done():
        http.Error(w, "Request cancelled", http.StatusRequestTimeout)
    }
}
```

## **Gin 框架**

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    r.Run(":8080")
}
```

### 获取请求参数

#### 1. URL 路径参数

```go
// GET /users/:id
r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"user_id": id})
})

// GET /users/:id/posts/:postId
r.GET("/users/:id/posts/:postId", func(c *gin.Context) {
    userId := c.Param("id")
    postId := c.Param("postId")
    c.JSON(200, gin.H{
        "user_id": userId,
        "post_id": postId,
    })
})
```

#### 2. Query 参数（URL ?key=value）

```go
// GET /search?q=golang&page=1
r.GET("/search", func(c *gin.Context) {
    // 获取单个参数
    q := c.Query("q")                    // 如果不存在返回 ""
    page := c.DefaultQuery("page", "1")  // 带默认值

    // 获取参数并检查是否存在
    sort, exists := c.GetQuery("sort")
    if exists {
        // sort 参数存在
    }

    c.JSON(200, gin.H{
        "query": q,
        "page":  page,
        "sort":  sort,
    })
})
```

#### 3. POST 表单数据

```go
// POST /form
r.POST("/form", func(c *gin.Context) {
    username := c.PostForm("username")
    password := c.DefaultPostForm("password", "default")

    // 检查是否存在
    email, exists := c.GetPostForm("email")

    c.JSON(200, gin.H{
        "username": username,
        "password": password,
        "email":    email,
        "exists":   exists,
    })
})
```

#### 4. JSON 请求体

```go
type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=120"`
}

// POST /users
r.POST("/users", func(c *gin.Context) {
    var user User

    // 绑定并验证 JSON
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "message": "User created",
        "user":    user,
    })
})
```

#### 5. 绑定多种格式

```go
type Login struct {
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

r.POST("/login", func(c *gin.Context) {
    var login Login

    // 自动根据 Content-Type 绑定
    if err := c.ShouldBind(&login); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"user": login.User})
})
```

### 响应方法

#### 1. JSON 响应

```go
// 返回 JSON
r.GET("/json", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "success",
        "data":    []int{1, 2, 3},
    })
})

// 返回结构体
type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

r.GET("/struct", func(c *gin.Context) {
    c.JSON(200, Response{
        Code:    200,
        Message: "OK",
    })
})

// 格式化的 JSON（带缩进）
r.GET("/json-pretty", func(c *gin.Context) {
    c.IndentedJSON(200, gin.H{"message": "pretty"})
})
```

#### 2. XML 响应

```go
r.GET("/xml", func(c *gin.Context) {
    c.XML(200, gin.H{"message": "success"})
})
```

#### 3. 字符串响应

```go
r.GET("/string", func(c *gin.Context) {
    c.String(200, "Hello %s", "World")
})
```

#### 4. HTML 响应

```go
r.LoadHTMLGlob("templates/*")

r.GET("/index", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "title": "Home Page",
    })
})
```

#### 5. 文件响应

```go
// 返回文件
r.GET("/file", func(c *gin.Context) {
    c.File("./static/image.png")
})

// 文件下载
r.GET("/download", func(c *gin.Context) {
    c.FileAttachment("./files/report.pdf", "report.pdf")
})

// 返回文件内容（从 io.Reader）
r.GET("/data", func(c *gin.Context) {
    data := []byte("file content")
    c.Data(200, "application/octet-stream", data)
})
```

#### 6. 重定向

```go
// HTTP 重定向
r.GET("/redirect", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "https://google.com")
})

// 路由重定向
r.GET("/old", func(c *gin.Context) {
    c.Request.URL.Path = "/new"
    r.HandleContext(c)
})
```

### 请求头操作

#### 读取请求头

```go
r.GET("/headers", func(c *gin.Context) {
    contentType := c.GetHeader("Content-Type")
    userAgent := c.GetHeader("User-Agent")
    auth := c.GetHeader("Authorization")

    // 获取所有请求头
    headers := c.Request.Header

    c.JSON(200, gin.H{
        "content-type": contentType,
        "user-agent":   userAgent,
    })
})
```

#### 设置响应头

```go
r.GET("/set-header", func(c *gin.Context) {
    c.Header("X-Custom-Header", "value")
    c.Header("Content-Type", "application/json")
    c.JSON(200, gin.H{"message": "ok"})
})
```

### Cookie 操作

```go
// 设置 Cookie
r.GET("/set-cookie", func(c *gin.Context) {
    c.SetCookie(
        "session_id",           // name
        "abc123",               // value
        3600,                   // maxAge (秒)
        "/",                    // path
        "localhost",            // domain
        false,                  // secure
        true,                   // httpOnly
    )
    c.String(200, "Cookie set")
})

// 读取 Cookie
r.GET("/get-cookie", func(c *gin.Context) {
    cookie, err := c.Cookie("session_id")
    if err != nil {
        c.JSON(400, gin.H{"error": "Cookie not found"})
        return
    }
    c.JSON(200, gin.H{"session_id": cookie})
})
```

### 文件上传

#### 单文件上传

```go
r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // 保存文件
    dst := "./uploads/" + file.Filename
    if err := c.SaveUploadedFile(file, dst); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "filename": file.Filename,
        "size":     file.Size,
    })
})
```

#### 多文件上传

```go
r.POST("/upload-multiple", func(c *gin.Context) {
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    files := form.File["files"]

    for _, file := range files {
        dst := "./uploads/" + file.Filename
        c.SaveUploadedFile(file, dst)
    }

    c.JSON(200, gin.H{
        "uploaded": len(files),
    })
})
```

### Context 存储和传递数据

#### 存储数据

```go
// 在中间件中存储数据
r.Use(func(c *gin.Context) {
    c.Set("user_id", 123)
    c.Set("username", "alice")
    c.Next()
})

// 在 handler 中获取数据
r.GET("/profile", func(c *gin.Context) {
    // 获取并转换类型
    userId, exists := c.Get("user_id")
    if !exists {
        c.JSON(401, gin.H{"error": "Unauthorized"})
        return
    }

    // 类型断言
    id := userId.(int)
    username := c.GetString("username")

    c.JSON(200, gin.H{
        "user_id":  id,
        "username": username,
    })
})
```

#### 便捷的类型获取方法

```go
r.GET("/data", func(c *gin.Context) {
    // 预先设置的数据
    c.Set("string_val", "hello")
    c.Set("int_val", 42)
    c.Set("bool_val", true)

    // 获取不同类型
    str := c.GetString("string_val")
    num := c.GetInt("int_val")
    flag := c.GetBool("bool_val")
    duration := c.GetDuration("timeout")
    time := c.GetTime("created_at")

    c.JSON(200, gin.H{
        "string": str,
        "int":    num,
        "bool":   flag,
    })
})
```

### 中断请求处理

```go
// 使用 Abort
r.Use(func(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(401, gin.H{"error": "Unauthorized"})
        c.Abort() // 停止后续处理
        return
    }
    c.Next()
})

// AbortWithStatus
r.GET("/forbidden", func(c *gin.Context) {
    c.AbortWithStatus(403)
})

// AbortWithStatusJSON
r.GET("/error", func(c *gin.Context) {
    c.AbortWithStatusJSON(500, gin.H{
        "error": "Internal Server Error",
    })
})
```

### 完整示例：RESTful API

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

var users = []User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func main() {
    r := gin.Default()

    // 获取所有用户
    r.GET("/users", func(c *gin.Context) {
        c.JSON(200, users)
    })

    // 获取单个用户
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        for _, user := range users {
            if fmt.Sprint(user.ID) == id {
                c.JSON(200, user)
                return
            }
        }
        c.JSON(404, gin.H{"error": "User not found"})
    })

    // 创建用户
    r.POST("/users", func(c *gin.Context) {
        var newUser User
        if err := c.ShouldBindJSON(&newUser); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        newUser.ID = len(users) + 1
        users = append(users, newUser)
        c.JSON(201, newUser)
    })

    // 更新用户
    r.PUT("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updateUser User

        if err := c.ShouldBindJSON(&updateUser); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        for i, user := range users {
            if fmt.Sprint(user.ID) == id {
                updateUser.ID = user.ID
                users[i] = updateUser
                c.JSON(200, updateUser)
                return
            }
        }
        c.JSON(404, gin.H{"error": "User not found"})
    })

    // 删除用户
    r.DELETE("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, user := range users {
            if fmt.Sprint(user.ID) == id {
                users = append(users[:i], users[i+1:]...)
                c.JSON(200, gin.H{"message": "User deleted"})
                return
            }
        }
        c.JSON(404, gin.H{"error": "User not found"})
    })

    r.Run(":8080")
}
```

## HTTP 客户端 (Client)

### 1. 基本请求

```go
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    // 简单的 GET 请求
    resp, err := http.Get("https://api.github.com/users/octocat")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Body: %s\n", body)
}
```

### 2. 不同 HTTP 方法

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
    baseURL := "http://localhost:8080"

    // GET 请求
    getRequest(baseURL + "/users")

    // POST 请求
    user := User{Name: "Charlie", Email: "charlie@example.com"}
    postRequest(baseURL+"/users", user)

    // PUT 请求
    user.ID = 1
    user.Name = "Updated Charlie"
    putRequest(baseURL+"/users/1", user)

    // DELETE 请求
    deleteRequest(baseURL + "/users/1")
}

func getRequest(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("GET request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("GET %s: %s\n", url, body)
}

func postRequest(url string, data interface{}) {
    jsonData, _ := json.Marshal(data)

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("POST request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("POST %s: %s\n", url, body)
}

func putRequest(url string, data interface{}) {
    jsonData, _ := json.Marshal(data)

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("Error creating PUT request: %v", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("PUT request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("PUT %s: %s\n", url, body)
}

func deleteRequest(url string) {
    req, err := http.NewRequest(http.MethodDelete, url, nil)
    if err != nil {
        log.Printf("Error creating DELETE request: %v", err)
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("DELETE request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("DELETE %s: Status %s\n", url, resp.Status)
}
```

# Process

> [goroutine](https://github.com/geekfun007/intro/blob/cursor/go-channel-principles-and-practice-claude-4.5-sonnet-thinking-a106/PATTERNS.md)

```go
func work() {
    fmt.Println("Hello from goroutine")
}

func workWithWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done() // Counter--
    fmt.Println("Hello from goroutine")
}

func workWithChan(ch chan int) {
	fmt.Println("Hello from goroutine")
	ch <- 1
}

func main() {
    go work() // Goroutine 以非阻塞的方式执行，它们会随着程序（main goroutine）的结束而消亡。如何等待 goroutine 执行完成

    // 睡眠操作，实际项目中不推荐这种写法，因为我们并不知道 goroutine 多久会执行结束，等待时间过长会影响程序整体效率，等待时间过短又无法确保所有 goroutine 执行完成
    time.Sleep(1 * time.Second)

    // WaitGroup 是 GO 语言 Sync 库中提供的常用并发措施，WaitGroup 是 GO 语言 Sync 库中提供的常用并发措施
    var wg sync.WaitGroup

    wg.Add(1) // Counter++
	go workWithWaitGroup(&wg)
	wg.Wait() // Counter == 0 塞并等待所有的 goroutine 结束，即当需要等待的 goroutine 的个数为 0 时，等待结束

    // channel 是 goroutine 之间的通信的桥梁，默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步
    ch := make(chan int)
    go workWithChan(ch)
    <-ch

    ch := make(chan int)
    ch <- 1 // 如果 channel 的容量为空，亦即“unbuffered channel”，只有在发送和接收端都已经 Ready 的情况下才会保证通信成功（否则会一直阻塞）
    <-ch

    ch := make(chan int, 1) // 如果 channel 的容量为空，亦即“unbuffered channel”，只有在发送和接收端都已经 Ready 的情况下才会保证通信成功（否则会一直阻塞）
	ch <- 1 // 向 channel 中传递一个 int 值，此时并不会阻塞程序
	<-ch

    close(ch) // close方法可以关闭 channel，channel 关闭后不可以再接收数据，如果继续向 channel 中发送数据的话会报错：panic: send on closed channel
    ch <- 2 // 向已经关闭的 channel 中发送数据会报错：panic: send on closed channel
    <-ch // 即使已经关闭了 channel，仍然可以从 channel 中读取 0

    for v := range ch { // 循环
        fmt.Println(v)
    }

    func onlyReceive(c <-chan string) { // 仅支持接受
        fmt.Println(<-c)
    }
    func onlySend(c chan<- string) { // 仅支持发送
        c <- "Send Only"
    }

    ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
        close(ch1)
     }()
	go func() {
        close(ch2)
     }()

	select { // Promise.race
	case <-ch1:
		fmt.Println("收到 ch1")
	case <-ch2:
		fmt.Println("收到 ch2")
	}

    // Context 传递共享数据
    type Context interface {
        Deadline() (deadline time.Time, ok bool)           // 获取截止时间
        Done() <-chan struct{}                             // 取消信号 channel
        Err() error                                        // 检查为何被取消（Canceled 或 DeadlineExceeded）
        Value(key interface{}) interface{}                 // 跨 API 传递只读数据
    }

    func Background() Context  // Background 通常作为顶级的Context用在main方法，Background 返回一个空的Context。它永远不会被取消，也没有超时，并且不携带任何值

    func WithCancel(parent Context) (ctx Context, cancel CancelFunc) // 当父Context的Done通道关闭或cancel方法被调用时，该拷贝的Done通道被关闭
    func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) // 当到达截止时间，或调用取消函数，或父context的完成通道关闭了，context拷贝完成通道也随之关闭
    func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) // 该拷贝的超时期限为now + timeout或父Context的超时期限，当该拷贝被取消时，会回收timer所占用的资源


}

func WithValue(parent Context, key, val interface{}) Context // 返回携带的传入key-value的父Context的拷贝

// 设置值
ctx := context.WithValue(parentCtx, key, value)

// 获取值
val := ctx.Value(key)
if val != nil {
    // 类型断言
    actualValue := val.(string)
}

// 如何修改 Context 中的 Value
// 为什么不能修改？Context 的不可变性是设计理念：并发安全: 多个 goroutine 可以安全地共享同一个 Context / 清晰的数据流: 避免隐式的状态变更 / 防止副作用: 子调用不会影响父调用的 Context

// 方案 1: 创建新的 Context（推荐）
package main

import (
    "context"
    "fmt"
)

type User struct {
    ID    string
    Name  string
    Score int
}

// 不推荐: 使用 string 作为 key, 容易冲突
// 推荐: 使用私有类型作为 key
type userKey struct{}

func main() {
    ctx := context.Background()

    // 初始用户
    user1 := &User{ID: "123", Name: "Alice", Score: 100}
    ctx = context.WithValue(ctx, userKey{}, user1)

    fmt.Printf("原始用户: %+v\n", ctx.Value(userKey{}))

    // "修改" - 实际是创建新的 Context
    user2 := &User{ID: "123", Name: "Alice", Score: 200}
    ctx = context.WithValue(ctx, userKey{}, user2)

    fmt.Printf("新用户: %+v\n", ctx.Value(userKey{}))
}

// 方案 2: 存储指针，修改指针指向的对象
type User struct {
	ID    string
	Name  string
	Score int
	mu    sync.RWMutex // 保护并发访问
}

// 线程安全的更新方法
func (u *User) UpdateScore(newScore int) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Score = newScore
}

// 线程安全的读取方法
func (u *User) GetScore() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Score
}

type userKey struct{}

func WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func GetUser(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(userKey{}).(*User)
	return user, ok
}

func demoPointerModification() {
	fmt.Println("=== 方案 2: 使用指针修改 ===")

	ctx := context.Background()

	user := &User{ID: "123", Name: "Alice", Score: 100}
	ctx = WithUser(ctx, user)

    if u, ok := GetUser(ctx); ok {
	    fmt.Printf("初始分数: %d\n", u.GetScore())
    }

    user.UpdateScore(200)
}
```
