# Go variable define

### Go variable type
- 

在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。

```go
import "fmt"
import "os"

const i = 100
const pi = 3.1415
const prefix = "Go_"

var i int
var pi float32
var prefix string
```

可以分组写成如下形式：


```go
import(
	"fmt"
	"os"
)

const(
	i = 100
	pi = 3.1415
	prefix = "Go_"
)

var(
	i int
	pi float32
	prefix string
)
```

_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：

```
_, b := 34, 35
```


## Go程序设计的一些规则
Go之所以会那么简洁，是因为它有一些默认的行为：

- 大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
- 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。




### reference
- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
- https://www.tutorialspoint.com/go/go_data_types.htm
