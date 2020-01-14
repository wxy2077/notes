package main

import (
	"../extends"
	"../util" // path
	"fmt"
)

func main() {
	const aaa = 123
	fmt.Println(aaa)
	fmt.Println(extends.Hello)
	fmt.Println(util2.Name) // package name
	fmt.Println(util2.Zzz)
}
