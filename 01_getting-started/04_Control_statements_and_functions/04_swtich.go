/*
Python写多了， 对switch很不习惯，感觉好多余， 哈哈哈哈

switch

有些时候你需要写很多的if-else来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，
而且也不易于以后的维护，这个时候switch就能很好的解决这个问题。它的语法如下


switch sExpr {
case expr1:
	some instructions
case expr2:
	some other instructions
case expr3:
	some other instructions
default:
	other code
}

sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，
表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；

而如果switch没有表达式，它会匹配true。

*/

package main

import "fmt"

func main() {
	myAAA()
}

func myAAA() {
	i := 1
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Printf("2")
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("以上都不是", i)
	}
}
