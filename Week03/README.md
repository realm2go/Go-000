### 2 detecting Race Conditions whith go

查看汇编 go tool compile -S fileName.go

Go memory model中提到：写入单个machine word是原子的
但interface 中是两个machine word的值，另外goroutine在更改interface的值时，可能出现问题。

假如内存布局不一样，可能出问题。

### 3 sysc.atomic