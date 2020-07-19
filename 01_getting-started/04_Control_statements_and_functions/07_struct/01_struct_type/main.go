/*

Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。
例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。如下代码所示:


*/
package _1_struct_type

import "fmt"

// 声明一个新的类型
type person struct {
	name string
	age  int
}

// 定义一个函数，比较年龄大小 返回年龄差
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person // P现在就是person类型的变量了

	tom.name = "Alice"                                 // 赋值"Alice"给P的name属性.
	tom.age = 25                                       // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s \n", tom.name) // 访问P的name属性.

	/*
		除了上面这种P的声明使用之外，还有另外几种声明使用方式：

		1.按照顺序提供初始化值

		P := person{"Tom", 25}

		2.通过field:value的方式初始化，这样可以任意顺序

		P := person{age:24, name:"Tom"}

		3.当然也可以通过new函数分配一个指针，此处P的类型为*person

		P := new(person)
	*/

	eric := person{"Eric", 15}

	who, diff := older(eric, tom)

	fmt.Println("person", who, "diff", diff)

}
