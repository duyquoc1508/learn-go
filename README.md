### Hot reload

with `nodemon` or `gin`,...

- With `nodemon`

```shell
nodemon --exec go run . --ext go
```

## Start new project

```
go mod init example.com/demo
```

### Add the Go driver as a dependency.

```
go get ...
```

### Functions

- log.Fatal(message)
- panic(err)
- defer cancel()

`:=`: Khai báo và gán giá trị
`=`: Gán giá trị, biến đã được khai báo trước đó

- Cú pháp này được gọi là One Liner if...else. If with short statement. If statement can start with a short statement to exec before condition

```go
if err := doStuff(); err != nil {
  // handle the error here
}
```

Giải thích cú pháp: Mệnh đề đầu tiên như là mệnh đề khởi tạo, nó hữu ích để khởi tạo biến cục bộ, chỉ sử dụng được trong block if đó

### Package

- By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package `rand`.
  -> `rand.Intn(231)`, `rand.Seed`

### Exported name

In Go, a name is exported if it begins with a Capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the math package.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

### Function

- Khi nhiều tham số hàm gần nhau chung 1 kiểu, ta có thể viết theo cách rút gọn

```go
func long(x int, y int, z string) - new -> func short(x, y int, z string)
```

- Golang allow function return multiple results

- Có 2 cách return giá trị trong hàm

  - return a, b -> return giá trị được chỉ định `func hello(x, y int) (string, string)`
  - return -> Khai báo giá trị biến được return ngay đầu hàm `func hello(x, y int) (z string, k string)` -> z, k lúc này đã được khai báo, trong hàm chỉ cần gán giá trị. cuối hàm không cần chỉ định `return z, k` mà chỉ cần `return`

- Short variable declarations (`:=`) only available inside function

### Basic type

`byte` alias for unit8

- Các biến được khai báo mà không gán giá trị mặc định thì go sẽ tự gán các giá trị mặc định tùy theo data type như sau

* numeric type - default: 0
* string type - default: ""
* boolean type - default: false

### Type conversions (ép kiểu)

The expression `T(v)` converts the value `v` to the type `T`.

### Loop

Go has only one looping construct, the `for` has 3 components separated by semicolons

- The init statement: Execute before the first iteration. <Optional>
- The condition statement: Evaluate before every iteration
- Post statement: Execute at the end of every iteration. <Optional>

* Có thể thiếu init statement và post statement. lúc này `for` hoạt động trong như `while`

### Defer (Hoãn lại)

A defer statement defers the execution of a function until the surrounding function returns. Được thực thi sau khi các hàm xung quanh thực thi xong
`defer fmt.Println("world")`

Deferred function calls are pushed onto a `stack`. When a function returns, its deferred calls are executed in `last-in-first-out` order

```go
package main
import "fmt"

func main() {
	defer fmt.Println("Defer in main")
	fmt.Println("Stat main")
	f1()
	fmt.Println("Finish main")
}

func f1() {
	defer fmt.Println("Defer in f1")
	fmt.Println("Start f1")
	f2()
	fmt.Println("Finish f1")
}

func f2() {
	defer fmt.Println("Defer in f2")
	fmt.Println("Start f2")
	fmt.Println("Finish f2")
}
```

Output

```txt
Stat main
Start f1
Start f2
Finish f2
Defer in f2
Finish f1
Defer in f1
Finish main
Defer in main
```

### Pointer

Pointer hold the memory address of a value

- The & operator generates a pointer to its operand. Toán tử & tạo ra 1 con trỏ đến toán hạng của nó

- The _ operator denotes the pointer's underlying value. Toán tử _ biểu thị giá trị của con trỏ

- Mặc định phép gán cho type struct trong golang là tham trị. Muốn tham chiếu thì phải dùng con trỏ. Sau đó thay đổi giá trị thì dùng hoàn toàn như js

```go
v := Vertex{1, 2}
shallowClone := &v
```

**Important**: Pointer trong function vs. method

- Function: hàm có argument là **pointer** `func ScaleFunc(v *Vertex, f float64)` thì thì đối số nhận vào (parameter) phải là 1 con trỏ (&v)

- Method: Phương thức của 1 type là **pointer** `func (v *Vertex) Scale(f float64)` thì biến gọi phương thức đó không cần phải là con trỏ. Go sẽ ngầm hiểu `v.Scale(5)` là `(&v).Scale`

Lợi ích của việc sử dụng con trỏ

- Function/method có thể thay đổi giá trị của biến

- Tránh copy giá trị mỗi lần function/method call. Có hiệu quả với param là 1 struct lớn

### Array

- Fixed sized

- Cùng 1 kiểu dữ liệu

- Array trong go không phải là dạng tham chiếu mà là tham trị. cần sử dụng list là tham số cho 1 hàm người ta đề xuất sử dụng slice thay vì array

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

### Slice

**Slice là 1 tham chiếu đến array, nó mô tả 1 phần hay toàn bộ array**

Dynamically-sized - allow select a half open range -> `slice[1:4]`

Nếu không có low bound -> mảng không bị cắt. nếu có low bound thì cap của slice sẽ bị giảm

Slice `[]bool{true, true, false}` -references-> Array `[n]bool{true, true, false}`

Slice có chiều dài (len) và sức chứa (cap)

Slice rỗng `var s []int` có len = 0 và cap = 0. thì mảng đó = nil. => `s == nill` = true

Slice can create by `make` function

`Length` là số phần tử chứa trong Slice

`Capacity` là số phần tử chứa trong Array mà Slice tham chiếu đến

Các cách để khai báo 1 slice

```go
c := []int{2, 3, 5, 7, 11, 13} // khai báo 1 mảng mà không chỉ ra kích thước thì đó là slice

a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

Nguyên lý hoạt động của cap trong slice

- Khi append 1 số lượng phân tử vào slice. Nếu cap của slice hiện tại không chứa đủ, một mảng lớn hơn sẽ được cấp phát có cap mới gấp <b>\*N</b> (`newCap = oldCap * N`) lần so với cap cũ. Tùy thuộc vào chiều dài mới mà N có thể là 1, 2, 3 ...

Ví dụ:

```go
var s []int = []int{1, 2, 3} // có len = 3 và cap = 3
append(s, 4) // lúc này len = 4 và cap = 6. Vì mảng cũ không đủ cap để chứa phần tử mới nên 1 mảng mới sẽ được cấp phát có cap * 2 so với mảng cũ
append(s, 5) // len = 5, cap = 6. Chỉ khi nào số phần tử thêm vào lớn hơn cap - len thì mảng mới mới được cấp phát
append(s, 6, 7) // len = 7, cap = oldCap * 2 = 12
```

### Map

```go
m := make(map[string]int) // declare map with make func
fmt.Println()

m["Answer"] = 42
fmt.Println("The value:", m["Answer"]) // The value: 42

m["Answer"] = 48
fmt.Println("The value:", m["Answer"]) // The value: 42

delete(m, "Answer") // delete a key
fmt.Println("The value:", m["Answer"]) // The value: 0. If key not exist in map

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok) // If key is in m, ok is true. If not, ok is false.
```

### Function

Function trong golang cũng có tính closures giống như js

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```

### Method

Go does not have classes. However, you can define methods on `types`

A method is a function with a special `receiver` argument

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

syntax

```go
func (t Type) methodName(params) returns {}  // method syntax
```

`(t Type)` là receiver có thể là **struct** hoặc **non-struct**
Có 2 loại receiver

- Value receiver: Không làm thay đổi giá trị receiver khi ra khỏi hàm

- Pointer receivers: This means the `receiver` type has the literal syntax \*T for some type T. Giá trị của receiver bị thay đôi nếu trong hàm có thay đổi

Function và method có agrument đều là tham trị (!= pointer)

- Function: Tham số khi gọi phải là giá trị (!= pointer). Function không cần thuộc về 1 đối tượng cụ thể
- Method: Type gọi method có thể là giá trị hoặc con trỏ. Nếu là con trỏ, go sẽ tự hiểu `p.Abs()` như là `(*p).Abs()`. Method thì phải thuốc 1 đối tượng nào đó (receiver)

### Interface

Là một tập các method mà một object có thể implement. Nó định nghĩa các hành vi của đối tượng. Thể hiện tính đa hình của đối tượng (Đều là hành vi `speak()` của lớp Animal. Nhưng mỗi đối tượng con sẽ `speak()` khác nhau. Ví dụ như Cat, Dog)

```go
type T struct {
  S string
}

type abcInterface interface {
  method1()
}

func (t *T) method1() { // function implement
}

func main() {
  var i abcInterface
  i.method1()
}
```

Trong golang, để implement 1 method trong interface ta chỉ cần đặt tên method của struct trùng tên với method trong interface mà chúng ta muốn implement

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods. Include any method

Empty interface. The interface type that specifies zero methods is known as the empty interface. An empty interface may hold values of **any type**. So that are used by code **handles value of unknown type**.

```go
func main() {
	var i interface{}
	describe(i) // (<nill>,<nill>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
  // or
  do(21)
	do("hello")
	do(true)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {}
```

Interface with init data (type assertions)

multiple interface: 1 struct có thể implement nhiều interface

Embedded interface: Có thể nhét nhiều interface vào 1 interface

```go
var i interface{} = "hello"
```

### Struct

- Named struct and anonymous struct. Trong go có 2 kiểu khai báo struct. Named struct là khai báo type struct và nơi gán giá trị cho struct riêng. Anonymous struct là cách khai báo struct không cần đặt tên cho struct, cách khai báo này buộc khởi tạo và gán giá trị 1 lần

- Named field and anonymout field. Named field là khai báo type struct với key - value. Anonymous field là khai báo struct không có key mà chỉ có value là data type của field

- Pointer in struct

- Nested struct

- Có 2 kiểu khởi tạo struct là: Khởi tạo tường mình - Khởi tạo không tường minh. Khởi tạo tường minh là khi khởi tạo có truyền key - value. Khởi tạo không tường minh là khi khởi tạo chỉ có truyền value theo thứ tự định nghĩa của struct mà không truyền key

- Có thể so sánh 2 struct. Chỉ cần 2 struct có các kiểu dữ liệu có thể só sánh được thì có thể so sánh được với nhau. Golang so sánh 2 struct theo kiểu tham trị (So sánh từng field của 2 struct)

- Viết hoa chữ cái đầu tiên của 1 `field name` của `struct` chứng tỏ field đó là public, viết thường là private

### Concurrency trong golang

Golang không gọi các thread trong concurrency mà với 1 tên gọi khác là `Goroutine`. Giao tiếp giữa các goroutine là `channel`

Khai báo 1 goroutine

```go
go doSomeWork()
```

Cú pháp làm việc với channel

```go
c <- data // đẩy dữ liệu vào channel
data = <- c // lấy dữ liệu từ channel ra
```

Trong go giao tiếp giữa các goroutine không trực tiếp share vùng nhớ mà sẽ giao tiếp qua channel

### Context <package>

`context.WithTimeout(parentContext, duration)`
`context.Background()`:
`context.TODO()`: Dùng khi chưa rõ context nào được áp dụng cho logic đó

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
```

Passing data to context with key is struct or other type

```go
type myStruct struct {
    ID  string
    Sig string
}
func main(){

	var ctx context.Context
	mySig := myStruct{
		ID:  "12345678",
		Sig: "Secret_Signature_Token",
	}

// Setting a Value associated with a Key in Context
ctx1 := context.WithValue(ctx, "myKey", &mySig)

//Getting the same value
	value, ok := ctx1.Value("myKey").(*myStruct) // đọc dữ liệu ở context và ép kiểu
	if ok {
	fmt.Println(value.ID)
	}
}
```

`ctx.Value("myKey")` return an interface. `.(*myStruct)` to convert it to the type `*myStruct`. So the value in the left hand side is of type `*myStruct` and u can access its field e.g `value.Sig`, `value.ID`

### bson <package>

format giá trị truyền vào để convert qua bson
`bson.M`: Kiểu map
`bson.D`: Kiểu slice. Mỗi phần tử là 1 struct
`bson.A`: Kiểu array
Ví dụ

```go
bson.M{"user": user} // kiểu map. Ngầm hiểu là {user : user}
bson.D{{ Key: "email", Value: "abc@gmail.com" }, { Key: "password", Value"123" }} // kiểu slice. mỗi phần tử là 1 struct. Ngầm hiểu là{email : "abc@gmai.com", password : "123"}
bson.A{"a", "b"} // kiểu array. Ngầm hiểu là {...: "a", ...: "b"}
```

### Marshal - Unmarshal

Marshal: Passing từ struct sang JSON. Dựa vào struct tag
Unmarshal: Passing từ JSON to struct

- Parsing JSON mà không biết trước cấu trúc. việc định nghĩa struct là bất khả thi do đó chúng ta sẽ dùng empty interface. Đợi đến thời điểm runtime thì compiler sẽ cung cấp memory phù hợp có những thứ đó

```go
var parsed interface{}
err := json.Unmarshal(data, &parsed)
```

### Goroutines

Goroutines là các hàm hoặc phương thức chạy đồng thời với các hàm/ phương thức khác

Trong go lúc nào cũng có 1 goroutine chính hay còn gọi là main goroutine. Khi Goroutine chính chạy xong mà các goroutines khác chưa chạy xong thì các goroutines khác đều bị hủy

Example

```go
package main

import (
    "fmt"
    "time"
)

func numbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```

Kết quả

```txt
1 a 2 3 b 4 c 5 d e main terminated
```

Hình dưới đây mô tả cách chương trình làm việc
![goroutines-explained](./images/Goroutines-explained.png 'goroutines explained')

### Channel
Cacs Goroutines giao tiếp với nhau bằng cách sử dụng channel. Tương tự như cách nước chảy từ đầu này sang đầu kia trong đường ống, dữ liệu có thể được gửi từ một đầu và nhận từ đầu kia bằng channels.

Gửi và nhận dữ liệu từ channel. Chú ý chiều mũi tên từ channel
```go
a := make(chan int)
data := <- a // đọc từ kênh a
a <- data // gửi từ kênh a
```

Gửi và nhận đến một channel đang bị chặn theo mặc định.
- Khi dữ liệu được **gửi** đến một channel, điều khiển sẽ bị chặn trong câu lệnh gửi cho đến khi một số Goroutine khác đọc từ channel đó.
- Tương tự khi dữ liệu được **đọc** từ một channel, việc đọc bị chặn cho đến khi một số Goroutine ghi dữ liệu vào channel đó.

Có thể tạo các channel 1 chiều. chỉ gửi hoặc chỉ nhận dữ liệu

Channel có thể có số lượng phần tử chưa trong channel đó gọi là (buffered channel). Khi các goroutine cố găng ghi dữ liệu vào channel lớn hơn số lượng khai báo thì đều bị chặn lại cho đến khi giá trị có sẵn bên trong được đọc
```go
ch := make(chan bool, 1)
```
### Select
Câu lệnh `select` được sử dụng để chọn từ nhiều hoạt động kênh gửi / nhận. Câu lệnh select sẽ chặn cho đến khi một trong các hoạt động gửi / nhận đã sẵn sàng. Nếu nhiều case đã sẵn sàng cùng 1 lúc, một trong số chúng được chọn ngẫu nhiên

Câu lệnh `select` sẽ chặn chương trình cho đến khi 1 trong các case của nó được thực thi

Cú pháp của select tương tự như switch. cũng có `default` case khi không có case nào khác sẵn sàng
```go
package main

import (
    "fmt"
    "time"
)

func server1(ch chan string) {
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}
func main() {
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
			case s1 := <-output1:
					fmt.Println(s1)
			case s2 := <-output2:
					fmt.Println(s2)
    }
}
```
Kết quả
```txt
from server2  // 2 goroutines đang chạy trong 1 block select. server2 chạy nhanh hơn nên server2 được chọn và in ra.
```

### Mutex
Mutex được sử dụng để cung cấp cơ chế khóa để đảm bảo rằng chỉ có một Goroutine đang chạy đoạn mã quan trọng tại bất kỳ thời điểm nào để ngăn chặn các xử lý ngoài mong đợi. Nhiều goroutines cùng thay đổi giá trị của 1 biến cũng 1 thời điểm có thể dẫn tới kết quả ngoài mong đợi (race condition)


## About this project

This project use `repository pattern`. The repository pattern included 3 component
![repository-pattern](./images/repository-pattern.png 'Repository pattern')
Có 3 thành phần chính:

- Interface: Chứa các method mà tương tác với table
- Interface implement: Implement các method đã khai báo trong interface. Trong file này chứa các phần code làm việc trức tiếp với database
- Model: Chức các file đại diện cho các table của database

Project tree

```bash
|__ utils
|   |__ response.util.go
|__ configs
    |__ index.config.go
|__ drivers
|  |__ <database>_driver.go
|__ main.go
|__ models
|  |__ user.model.go
|__ repositories
   |__ repoImpl
   |   |__ user.repoImpl.go
   |__ user.repo.go
```

## Ref

- https://go.dev/tour
