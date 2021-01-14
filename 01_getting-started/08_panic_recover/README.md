
参考
- https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/
- https://hezhiqiang8909.gitbook.io/go/ch1-basic/ch1-07-error-and-panic

> Go中大量的 err != nil 给人感觉很不好。

## panic 触发的递归延迟调用

- `panic` 能够改变程序的控制流，调用 `panic` 后会立刻停止执行当前函数的剩余代码，并在当前 `Goroutine` 中递归执行调用方的 `defer`；
- `recover` 可以中止 `panic 造成的程序崩溃。它是一个只能在 `defer` 中发挥作用的函数，在其他作用域中调用不会发挥作用；

我们先通过几个例子了解一下使用 `panic` 和 `recover` 关键字时遇到的现象，部分现象也与上一节分析的 `defer` 关键字有关：

- `panic` 只会触发当前 `Goroutine` 的 `defer`；
- `recover` 只有在 `defer` 中调用才会生效；
- `panic` 允许在 `defer` 中嵌套多次调用；

