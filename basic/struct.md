# struct 类型


**1. 可以通过 struct 声明新的类型。**
```go

type Person struct {
	name string
	age int
}

var p Person

p.name = "Json Mic"
p.age = 18

// 根据顺序赋值
p1 := Person{"Json Mic", 18}
// 根据 key 赋值
p2 := Person{name: "Json Mic", age: 18}

// new 分配指针
p3 := new(Person)
```

**2. 继承另一个类型的属性**

```go
type Person struct {
	name string
	age int
}

// User 隐式的扩展了 Person 的属性
type User {
	Person
	address string
	phone string
}

var u User = User{Person{"Tyler Swift", 18}, "American", "18765438971"}

u.name // Tyler Swift 
u.age // 18
u.address // American
u.Person.name // Tyler Swift
```

**注意：** 如果继承的属性值与本身属性值冲突，以本身的属性值为主。如果多个继承之间属性冲突，则属性无法继承。