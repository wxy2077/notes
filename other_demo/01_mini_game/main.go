package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

func main() {
	var num int
	var randomNum = getRandomNum()
	fmt.Printf("请输入数字:")
	fmt.Scanf("%d", &num)
	if num == 1 && randomNum == 2 {

	}

}
