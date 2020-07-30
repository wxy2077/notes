package main

import "fmt"

type Cat struct {
	name   string
	color  string
	sex    bool
	couple Mouse // 这里赋值为 另一个结构体, 如果和结构体名一样 可以省略一个
}

// 结构体挂载方法
func (c *Cat) CaptureMouse(name string) (res bool) {
	if c.name == name {
		fmt.Println("名字为Tom，能捉到老鼠")
		return true
	} else {
		fmt.Println("名字不为Tom， 捉不到老鼠")
		return false
	}
}

// 在定义一个结构体
type Mouse struct {
	name string
}

func (m Mouse) Eat(stuff string) (res bool) {
	if stuff == "奶酪" {
		fmt.Println(m.name, "喜欢吃", stuff)
		return true
	} else {
		fmt.Println(m.name, "不喜欢吃", stuff)
		return false
	}
}

func main() {
	// 实例化结构体
	tom := Cat{
		name:  "Tom",
		color: "white",
		sex:   true,
		couple: Mouse{
			name: "Jerry",
		},
	}

	// 调用方法
	res := tom.CaptureMouse("Tom")
	fmt.Println(res)

	tom.couple.Eat("糖")
	tom.couple.Eat("奶酪")

}
