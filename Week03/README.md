### 2 detecting Race Conditions whith go

查看汇编 go tool compile -S fileName.go

Go memory model中提到：写入单个machine word是原子的
但interface 中是两个machine word的值，另外goroutine在更改interface的值时，可能出现问题。

假如内存布局不一样，可能出问题。

### 3 sync.atomic

### 4 sync.mutex
几种互斥锁的实现：
- Baring   // 性能好，不公平
- Handoff  // 平衡性能和公平性
- Spining  // 自旋, 当等待队列为空,就没有必要快速释放锁,park 和 unpark也有成本

### 4 errgroup
核心原理:利用sync.Waitgroup管理并行执行的goroutine

### 5 sync.Pool

### 6 context

```
type Context interface {
    Deadline() (deadline time.Time,ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface) interface{}
}
```



> 有两种方法将context对象集成到API中:
- 首参数传context
- 在第一个request对象中携带一个可选的context对象

不要在结构体中传递context(要显示传递，不然调用者无法知道) 

源码：

1. context.WithValue(ctx,newvalue)

    - cow
    - 每次挂值，都需要繁殖生成一个新的context
    
2. context.Cancel()级联取消


    