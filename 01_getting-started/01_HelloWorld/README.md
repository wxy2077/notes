# GoLang program language getting started

<a href="./README_zh">中文</a>

> official documents address

- https://golang.google.cn


### go environment variable

The command ```go env GOPATH ``` print the effective current go path, it prints the default location if the environment variable is unset.

For convenience, add the workspace's bin subdirectory to your PATH:

```export PATH=$PATH:$(go env GOPATH)/bin```

-----
### First program
> https://golang.google.cn/doc/code.html#Overview

1 first step, you can create `hello.go` files

```
package main  // Defining package names.
// You must statement  the package name in the code first line (not notes) of the source files, e.g. package main,
// package main indicates a program that independent and enforceable, each go application should be included a package called main.
                                           
            

import "fmt"  // The full name of fmt package is format, which implements the function of formatting IO (input / output).

func main() {
	fmt.Println("Hello, world.") // Line feed printing
                                    
}
```

