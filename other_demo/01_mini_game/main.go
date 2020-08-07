package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机数
0为石头
1为剪刀
2为布
*/
func GetRandomNum() int {
	// 让随机数重置
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

/*
检测数据合法性
*/
func CheckInput(num int) bool {
	if num >= 0 && num <= 2 {
		return true
	} else {
		return false
	}
}

/*
转换对应的数字
*/
func TransNum(num int) string {
	switch {
	case num == 0:
		return "✊"
	case num == 1:
		return "✂️"
	default:
		return "🤚"

	}
}

func main() {

	// 整体无限循环
	for {
		var num int
		var randomNum = GetRandomNum()

		fmt.Printf("请输入数字【0:✊; 1:✂️; 2:🤚】:")

		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println("输入数据不合法", num)
			continue
		}
		res := CheckInput(num)
		if res == false {
			fmt.Println("输入数据应>=0并且<=2")
			continue
		}

		fmt.Println("电脑出的", TransNum(randomNum), " VS ", "你出的", TransNum(num))
		if (num == 0 && randomNum == 1) || (num == 1 && randomNum == 2) || (num == 2 && randomNum == 0) {
			fmt.Println("你赢了")
		} else if num == randomNum {
			fmt.Println("平局")
		} else {
			fmt.Println("电脑赢了")
		}
	}

}
