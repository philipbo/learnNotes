## context.context

**通过context，我们可以方便地对同一个请求所产生地goroutine进行约束管理，可以设定超时、deadline，甚至是取消这个请求相关的所有goroutine。形象地说，假如一个请求过来，需要A去做事情，而A让B去做一些事情，B让C去做一些事情，A、B、C是三个有关联的goroutine，那么问题来了：假如在A、B、C还在处理事情的时候请求被取消了，那么该如何优雅地同时关闭goroutine A、B、C呢？这个时候就轮到context包上场了。**

### context.Context 的使用规范
* 不要把context存储在结构体中，而是要显式地进行传递
* 把context作为第一个参数，并且一般都把变量命名为ctx
* 就算是程序允许，也不要传入一个nil的context，如果不知道是否要用context的话，用context.TODO()来替代
* context.WithValue()只用来传递请求范围的值，不要用它来传递可选参数
* 就算是被多个不同的goroutine使用，context也是安全的

### context.Context 结构
```
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this `Context` is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this Context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key interface{}, val interface{}) Context
```
### 继承 context

context 包提供了一些函数，协助用户从现有的 Context 对象创建新的 Context 对象。
这些 Context 对象形成一棵树：当一个 Context 对象被取消时，继承自它的所有 Context 都会被取消。

`Background` 是所有 Context 对象树的根，它不能被取消。它的声明如下：
```
// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level `Context` for incoming requests.
func Background() Contex
```

`WithCancel` 和 `WithTimeout` 函数会返回继承的 Context 对象， 这些对象可以比它们的父 Context 更早地取消。

当请求处理函数返回时，与该请求关联的 Context 会被取消。 当使用多个副本发送请求时，可以使用 WithCancel取消多余的请求。 WithTimeout 在设置对后端服务器请求截止时间时非常有用。 下面是这三个函数的声明：
```
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue` 函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：
```
// WithValue returns a copy of parent whose Value method returns val for key.
func WithValue(parent Context, key interface{}, val interface{}) Context
```