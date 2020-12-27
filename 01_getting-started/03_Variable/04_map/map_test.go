/*
map也就是Python中字典的概念，它的格式为map[keyType]valueType

我们看下面的代码，map的读取和设置也类似slice一样，通过key来操作，
只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。


*/
package _04_map

import (
	"testing"
)

func TestMap(t *testing.T) {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int // 还需要make

	//numbers := map[string]int{}  // 缩写后面加 {}
	//numbers := make(map[string]int)  // 或者直接 make
	// 另一种map的声明方式
	numbers2 := make(map[string]int)
	numbers2["one"] = 1  //赋值
	numbers2["ten"] = 10 //赋值
	numbers2["three"] = 3
	// 没有赋值  int类型默认为 0
	foo := numbers2["four"]
	t.Log(foo)
	//numbers2["five"] = 0
	// 如果赋值为0 如何判断是空值 还是赋的值
	if v, ok := numbers2["five"]; ok {
		t.Log("有five这个key的值", v)
	} else {
		t.Log("没有有five这个key的值", v)
	}

	t.Log("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3
	/*
		使用map过程中需要注意的几点：

		map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
		map的长度是不固定的，也就是和slice一样，也是一种引用类型
		内置的len函数同样适用于map，返回map拥有的key的数量
		map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
		map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

	*/

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	t.Log("初始化长度是: ", len(rating))
	delete(rating, "C")
	t.Log(len(rating))
	// 循环map k v  如果不想使用key或者value 可使用 _ 丢弃
	for k, v := range rating {
		t.Log(k, v)
	}
	t.Log("分割---------")

	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]

	if ok {
		t.Log("C# is in the map and its rating is ", csharpRating)
	} else {
		t.Log("We have no rating associated with C# in the map")
		t.Log("csharpRating", csharpRating)
	}

	delete(rating, "C") // 删除key为C的元素
	t.Log("---------------")

	//上面说过了，map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
	/*
		map[Hello:Salut]
		map[Hello:Salut]
	*/

	t.Log(m)
	t.Log(m1)
	// 定义 一个接口map
	m4 := make(map[string]interface{})
	m4["a"] = 123
	m4["b"] = "666"
	m4["c"] = true
	m4["d"] = []interface{}{1, "2", "3", true}

	t.Log(m4)

}
