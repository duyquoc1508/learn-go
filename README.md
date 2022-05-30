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

Deferred function calls are pushed onto a `stack`. When a function returns, its deferred calls are executed in `last-in-first-out` order. Defer đọc đầu tiên thì sẽ ra cuối cùng

### Pointer

Pointer hold the memory address of a value

- The & operator generates a pointer to its operand. Toán tử & tạo ra 1 con trỏ đến toán hạng của nó

- The _ operator denotes the pointer's underlying value. Toán tử _ biểu thị giá trị của con trỏ

- Mặc định phép gán cho type struct trong golang là tham trị. Muốn tham chiếu thì phải dùng con trỏ. Sau đó thay đổi giá trị thì dùng hoàn toàn như js

```go
v := Vertex{1, 2}
shallowClone := &v
```

### Array

Fixed sized

### Slice

Dynamically-sized - allow select a half open range -> `slice[1:4]`

Nếu không có low bound -> mảng không bị cắt. nếu có low bound thì cap của slice sẽ bị giảm

Slice thực ra làm tham chiếu đến 1 mảng

Slice `[]bool{true, true, false}` -references-> Array `[n]bool{true, true, false}`

Slice có chiều dài (len) và sức chứa (cap)

Slice rỗng `var s []int` có len = 0 và cap = 0. thì mảng đó = nil. => `s == nill` = true

Slice can create by `make` function

```go
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

## Ref

- https://go.dev/tour
