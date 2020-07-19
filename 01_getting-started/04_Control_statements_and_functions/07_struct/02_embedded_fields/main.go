/*

embedded  v 把…牢牢地嵌入(或插入、埋入);派遣(战地记者、摄影记者等);嵌入(在I'm aware that she knows句中，she knows为内嵌句)


struct的匿名字段  嵌入字段
我们上面介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，
而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

让我们来看一个例子，让上面说的这些更具体化
*/

package main

import "fmt"

type skill []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human // // 匿名字段，struct  嵌套更准确
	skill // 匿名字段，自定义的类型string slice
	//int  // 内置类型作为匿名字段
	speciality string
	hobby      []string
}

func main() {
	otis := Student{Human: Human{"Otis", 35, 100}, speciality: "Biology"}

	fmt.Println(otis)

	otis.Human.age = 15
	otis.hobby = []string{"read book"}

	otis.skill = []string{"cook"}

	fmt.Println(otis)

}
